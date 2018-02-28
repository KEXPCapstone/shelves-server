package sessions

import (
	"errors"
	"net/http"
)

const headerAuthorization = "Authorization"
const paramAuthorization = "auth"
const schemeBearer = "Bearer "

//ErrNoSessionID is used when no session ID was found in the Authorization header
var ErrNoSessionID = errors.New("no session ID found in " + headerAuthorization + " header")

//ErrInvalidScheme is used when the authorization scheme is not supported
var ErrInvalidScheme = errors.New("authorization scheme not supported")

//BeginSession creates a new SessionID, saves the `sessionState` to the store, adds an
//Authorization header to the response with the SessionID, and returns the new SessionID
func BeginSession(signingKey string, store Store, sessionState interface{}, w http.ResponseWriter) (SessionID, error) {
	sid, err := NewSessionID(signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	if err := store.Save(sid, sessionState); err != nil {
		return InvalidSessionID, err
	}
	w.Header().Add(headerAuthorization, schemeBearer+sid.String())
	return sid, nil
}

//GetSessionID extracts and validates the SessionID from the request headers
func GetSessionID(r *http.Request, signingKey string) (SessionID, error) {
	auth := r.Header.Get(headerAuthorization)
	if auth == "null" { // if client makes a fetch request with the Authorization header set to null, that is the actual string value of auth
		return InvalidSessionID, errors.New("not signed in")
	}
	if len(auth) == 0 {
		auth = r.URL.Query().Get("auth")
	}
	if len(auth) == 0 {
		return InvalidSessionID, errors.New("could not find auth query string parameter")
	}
	// TODO: strings.TrimPrefix instead
	auth = auth[len(schemeBearer):]
	sid, err := ValidateID(auth, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	return sid, nil
}

//GetState extracts the SessionID from the request,
//gets the associated state from the provided store into
//the `sessionState` parameter, and returns the SessionID
func GetState(r *http.Request, signingKey string, store Store, sessionState interface{}) (SessionID, error) {
	id, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	if err := store.Get(id, sessionState); err != nil {
		return InvalidSessionID, err
	}
	return id, nil
}

//EndSession extracts the SessionID from the request,
//and deletes the associated data in the provided store, returning
//the extracted SessionID.
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {
	id, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	if err := store.Delete(id); err != nil {
		return InvalidSessionID, err
	}
	return id, nil
}
