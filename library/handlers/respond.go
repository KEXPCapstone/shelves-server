package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respond(w http.ResponseWriter, statusCode int, value interface{}) {
	w.Header().Add(headerContentType, contentTypeJSON)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, fmt.Sprintf(ErrEncodingJSON+"%v", err), http.StatusInternalServerError)
		return
	}
}
