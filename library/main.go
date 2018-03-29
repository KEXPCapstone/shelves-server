package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	mux := http.NewServeMux()

	// TODO: Register handlers on mux

	log.Printf("Library microservice is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

// TODO: Implement handler function
