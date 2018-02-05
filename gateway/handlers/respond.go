package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// will write the provided http status code and value to the provided
// response writer.  Intended for use with values meant to be encoded
// as a JSON in an http response.
func respond(w http.ResponseWriter, statusCode int, value interface{}) {
	w.Header().Add(headerContentType, contentTypeJSON)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, fmt.Sprintf("Error converting response value to JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
