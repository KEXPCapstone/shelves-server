package handlers

import "net/http"

// /v1/shelves
func (hCtx *HandlerCtx) ShelvesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

// /v1/shelves/{id}
func (hCtx *HandlerCtx) ShelfHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}
