package handlers

import "net/http"

// middleware handler for supporting Cross Origin Requests
type CORSHandler struct {
	Handler http.Handler
}

// returns a new CORSHandler middleware handler
func NewCorsHandler(destHandler http.Handler) *CORSHandler {
	return &CORSHandler{destHandler}
}

// adds various headers to to response writer as defined by the constants.go file
func (ch *CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(accessControlAllowOrigin, valAllowOrigin)
	w.Header().Add(accessControlAllowMethods, valAllowMethods)
	w.Header().Add(accessControlAllowHeaders, valAllowHeaders)
	w.Header().Add(accessControlExposeHeaders, valExposeHeaders)
	w.Header().Add(accessControlMaxAge, valMaxAge)
	if r.Method != http.MethodOptions {
		ch.Handler.ServeHTTP(w, r)
	}
}
