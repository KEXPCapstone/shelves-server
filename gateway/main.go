package main

import (
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
	"github.com/go-redis/redis"
	mgo "gopkg.in/mgo.v2"
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
		// redisAddr = "localhost:6379"
		log.Fatal("Please provide REDISADDR")
	}

	dbAddr := os.Getenv("DBADDR")
	if len(dbAddr) == 0 {
		// dbAddr = "localhost:27017"
		log.Fatal("Please provide DBADDR")
	}

	// Commented out because of not being used yet
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

	// TODO: Unsure if we need repetition with the /
	mux.Handle("/v1/shelves", MicroServiceProxy(splitShelvesSvcAddrs, sessKey, rs))
	mux.Handle("/v1/shelves/", MicroServiceProxy(splitShelvesSvcAddrs, sessKey, rs))
	// TODO: May want structure it like /v1/library/release
	mux.Handle("/v1/library/releases", MicroServiceProxy(splitLibrarySvcAddrs, sessKey, rs))
	mux.Handle("/v1/library/releases/", MicroServiceProxy(splitLibrarySvcAddrs, sessKey, rs))

	corsHandler := handlers.NewCorsHandler(mux)
	fmt.Printf("'Gateway' server has been started at http://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCert, tlsKey, corsHandler)) // report if any errors occur
}

func MicroServiceProxy(addrs []string, signingKey string, sessStore sessions.Store) *httputil.ReverseProxy {
	index := 0
	mx := sync.Mutex{}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			// sessionState := &handlers.SessionState{}
			// _, err := sessions.GetState(r, signingKey, sessStore, sessionState)
			// if err == nil { // Add X-User header if signed in.
			// 	userJSON, jsonErr := json.Marshal(sessionState.AuthUsr)
			// 	if jsonErr != nil { // we know the user will be a json structured object, this error is unlikely to occur
			// 		log.Printf("error marshalling user: %v", sessionState.AuthUsr)
			// 	}
			// 	r.Header.Add("X-User", string(userJSON))
			// } else {
			// 	r.Header.Del("X-User") // remove Header in case hostile client tried to pass X-User
			// }
			mx.Lock()
			r.URL.Host = addrs[index%len(addrs)]
			index++
			mx.Unlock()
			r.URL.Scheme = "http" //downgrade to http protocol
		},
	}
}
