package main

import "net/http"

// TODO: Implement handler function
func (hCtx *HandlerCtx) ReleasesHandler(w http.ResponseWriter, r *http.Request) {
	// GetAll
	// GetByField
	// Insert
	field := r.URL.Query().Get("field")
	value := r.URL.Query().Get("value")

	switch r.Method {
	case http.MethodPost:
		// insert
	case http.MethodGet:
		if len(field) != 0 && len(value) != 0 {
			// mongo filter
		} else {
			// get all
		}
	}

}

func (hCtx *HandlerCtx) SingleReleaseHandler(w http.ResponseWriter, r *http.Request) {

}
