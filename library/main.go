package main

import (
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	mux := http.NewServeMux()

	dbAddr := os.Getenv("DBADDR")
	if len(dbAddr) == 0 {
		// dbAddr = "localhost:27017"
		log.Fatal("Please provide DBADDR")
	}

	mongoSess, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	mongoStore := NewMgoStore(mongoSess, "releasestore", "releases")

	// TODO: Register handlers on mux

	log.Printf("Library microservice is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

// TODO: Implement handler function
func ReleasesHandler(w http.ResponseWriter, r *http.Request) {
	// GetAll
	// Insert
}

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {

}

func GetByFieldHandler(w http.ResponseWriter, r *http.Request) {

}
