package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/KEXPCapstone/shelves-server/gateway/handlers"
	"github.com/KEXPCapstone/shelves-server/gateway/models/users"
	"github.com/KEXPCapstone/shelves-server/gateway/sessions"
	mgo "github.com/globalsign/mgo"
	"github.com/go-redis/redis"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKey := os.Getenv("TLSKEY")
	tlsCert := os.Getenv("TLSCERT")
	if len(tlsKey) == 0 || len(tlsCert) == 0 {
		log.Fatal("Please provide TLSKEY and TLSCERT")
	}

	sessKey := os.Getenv("SESSIONKEY")
	if len(sessKey) == 0 {
		log.Fatal("Please provide SESSIONKEY")
	}

	redisAddr := os.Getenv("REDISADDR")
	if len(redisAddr) == 0 {
		log.Fatal("Please provide REDISADDR")
	}

	dbAddr := os.Getenv("DBADDR")
	if len(dbAddr) == 0 {
		log.Fatal("Please provide DBADDR")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	rs := sessions.NewRedisStore(redisClient, time.Duration(10)*time.Minute)

	mongoSess, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	librarySvcAddr := os.Getenv("LIBRARYSVCADDR")
	if len(librarySvcAddr) == 0 {
		log.Fatal("Please provide LIBRARYSVCADDR")
	}
	splitLibrarySvcAddrs := strings.Split(librarySvcAddr, ",")

	shelvesSvcAddr := os.Getenv("SHELVESSVCADDR")
	if len(shelvesSvcAddr) == 0 {
		log.Fatal("Please provide SHELVESSVCADDR")
	}
	splitShelvesSvcAddrs := strings.Split(shelvesSvcAddr, ",")

	mongoStore := users.NewMgoStore(mongoSess, "userstore", "users")

	hCtx := handlers.NewHandlerContext(sessKey, rs, mongoStore)
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/users", hCtx.UsersHandler)
	mux.HandleFunc("/v1/users/me", hCtx.UsersMeHandler)
	mux.HandleFunc("/v1/sessions", hCtx.SessionsHandler)
	mux.HandleFunc("/v1/sessions/mine", hCtx.SessionsMineHandler)
	mux.Handle("/v1/shelves", MicroServiceProxy(splitShelvesSvcAddrs, sessKey, rs))
	mux.Handle("/v1/shelves/", MicroServiceProxy(splitShelvesSvcAddrs, sessKey, rs))
	mux.Handle("/v1/library/", MicroServiceProxy(splitLibrarySvcAddrs, sessKey, rs))
	
	corsHandler := handlers.NewCorsHandler(mux)
	log.Println("Starting to redirect http traffic to https")
	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
			log.Fatalf("HTTP ListenAndServe error: %v", err)
		}
	}()

	fmt.Printf("'Gateway' server has been started at http://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCert, tlsKey, corsHandler)) // report if any errors occur
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}

func MicroServiceProxy(addrs []string, signingKey string, sessStore sessions.Store) *httputil.ReverseProxy {
	index := 0
	mx := sync.Mutex{}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			sessionState := &handlers.SessionState{}
			_, err := sessions.GetState(r, signingKey, sessStore, sessionState)
			if err == nil { // Add X-User header if signed in.
				userJSON, jsonErr := json.Marshal(sessionState.AuthUsr)
				if jsonErr != nil {
					log.Printf("error marshalling user: %v", sessionState.AuthUsr)
				}
				r.Header.Add("X-User", string(userJSON))
			} else {
				r.Header.Del("X-User") // remove Header in case hostile client tried to pass X-User
			}
			mx.Lock()
			r.URL.Host = addrs[index%len(addrs)]
			index++
			mx.Unlock()
			r.URL.Scheme = "http" // downgrade to http protocol
		},
	}
}
