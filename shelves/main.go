package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KEXPCapstone/shelves-server/shelves/handlers"
	"github.com/KEXPCapstone/shelves-server/shelves/models"
	mgo "github.com/globalsign/mgo"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	mux := http.NewServeMux()

	dbAddr := os.Getenv("DBADDR")
	if len(dbAddr) == 0 {
		log.Fatal("Please provide DBADDR")
	}

	shelvesDb := os.Getenv("SHELVESDB")
	if len(shelvesDb) == 0 {
		log.Fatal("Please provide SHELVESDB")
	}

	shelvesColl := os.Getenv("SHELVESCOLL")
	if len(shelvesColl) == 0 {
		log.Fatal("Please provide SHELVESCOLL")
	}

	mongoSess, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	mongoStore := models.NewMgoStore(mongoSess, shelvesDb, shelvesColl)

	hCtx := handlers.NewHandlerContext(mongoStore)
	mux.HandleFunc("/v1/shelves", hCtx.ShelvesHandler)
	mux.HandleFunc("/v1/shelves/mine", hCtx.ShelvesMineHandler)
	mux.HandleFunc("/v1/shelves/users/", hCtx.UserShelvesHandler)
	mux.HandleFunc("/v1/shelves/featured", hCtx.FeaturedShelvesHandler)
	mux.HandleFunc("/v1/shelves/search/", hCtx.ShelvesForReleaseHandler)
	mux.HandleFunc("/v1/shelves/", hCtx.ShelfHandler)

	log.Printf("The 'shelves' microservice is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
