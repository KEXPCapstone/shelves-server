package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KEXPCapstone/shelves-server/library/handlers"
	"github.com/KEXPCapstone/shelves-server/library/models/releases"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
		// addr = "localhost:4001"
	}
	mux := http.NewServeMux()

	dbAddr := os.Getenv("DBADDR")
	if len(dbAddr) == 0 {
		// dbAddr = "localhost:27017"
		log.Fatal("Please provide DBADDR")
	}

	releaseDb := os.Getenv("RELEASEDB")
	if len(releaseDb) == 0 {
		log.Fatal("Please provide RELEASEDB")
	}
	releaseColl := os.Getenv("RELEASECOLL")
	if len(releaseColl) == 0 {
		log.Fatal("Please provide RELEASECOLL")
	}
	artistColl := os.Getenv("ARTISTCOLL")
	if len(artistColl) == 0 {
		log.Fatal("Please provide ARTISTCOLL")
	}
	genreColl := os.Getenv("GENRECOLL")
	if len(genreColl) == 0 {
		log.Fatal("Please provide GENRECOLL")
	}
	mongoSess, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	mongoStore := releases.NewMongoStore(mongoSess, releaseDb, releaseColl, artistColl, genreColl)

	releaseTrie, err := mongoStore.IndexReleases()
	if err != nil {
		// TODO
	}

	hCtx := handlers.NewHandlerContext(mongoStore, releaseTrie)

	mux.HandleFunc("/v1/library/releases", hCtx.ReleasesHandler)
	mux.HandleFunc("/v1/library/releases/", hCtx.SingleReleaseHandler)
	mux.HandleFunc("/v1/library/artists", hCtx.ArtistsHandler)
	mux.HandleFunc("/v1/library/genres", hCtx.GenresHandler)

	log.Printf("Library microservice is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
