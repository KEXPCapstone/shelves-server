package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
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

	// TODO: Microservice ADDRS

	// Commented out because of not being used yet
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	rs := sessions.NewRedisStore(redisClient, time.Duration(10)*time.Minute)

	mongoSess, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	mongoStore := users.NewMgoStore(mongoSess, "userstore", "users")

	// TODO: hCtx := handlers.NewHandlerContext()
	mux := http.NewServeMux()

	// TODO: Register handlers

	corsHandler := handlers.NewCorsHandler(mux)
	fmt.Printf("Server has been started at http://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCert, tlsKey, corsHandler)) // report if any errors occur
}

func MicroServiceProxy(addrs []string, signingKey string, sessStore sessions.Store) *httputil.ReverseProxy {
	// TODO`
	return nil
}
