package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KEXPCapstone/shelves-server/gateway/models/users"
	"github.com/KEXPCapstone/shelves-server/gateway/sessions"
)

func (hCtx *HandlerCtx) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		nu := &users.NewUser{}
		if err := json.NewDecoder(r.Body).Decode(nu); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON into new user: %v", err), http.StatusBadRequest)
			return
		}
		if err := nu.Validate(); err != nil {
			http.Error(w, fmt.Sprintf("Error creating new user: %v", err), http.StatusBadRequest)
			return
		}
		if _, err := hCtx.userStore.GetByEmail(nu.Email); err != users.ErrUserNotFound {
			http.Error(w, fmt.Sprintf("Another user already has the email address: %v", nu.Email), http.StatusBadRequest)
			return
		}
		if _, err := hCtx.userStore.GetByUserName(nu.UserName); err != users.ErrUserNotFound {
			http.Error(w, fmt.Sprintf("Another user already has the user name: %v", nu.UserName), http.StatusBadRequest)
			return
		}
		usr, err := hCtx.userStore.Insert(nu)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error adding user: %v", err), http.StatusInternalServerError)
			return
		}
		ss := &SessionState{
			AuthUsr:   usr,
			StartTime: time.Now(),
		}
		if _, err := sessions.BeginSession(hCtx.signingKey, hCtx.sessionStore, ss, w); err != nil {
			http.Error(w, fmt.Sprintf("Error initiating session: %v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusCreated, usr)
	default:
		http.Error(w, "Only allowed to POST to this resource", http.StatusMethodNotAllowed)
		return
	}
}

func (hCtx *HandlerCtx) UsersMeHandler(w http.ResponseWriter, r *http.Request) {
	ss := &SessionState{}
	_, err := sessions.GetState(r, hCtx.signingKey, hCtx.sessionStore, ss)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching current user: %v", err), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, http.StatusOK, ss.AuthUsr)
	default:
		http.Error(w, "Only allowed to GET to this resource", http.StatusMethodNotAllowed)
		return
	}
}

// TODO: To be revisited to assess intended authentication method
func (hCtx *HandlerCtx) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		cred := &users.Credentials{}
		if err := json.NewDecoder(r.Body).Decode(cred); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON of credentials: %v", err), http.StatusBadRequest)
			return
		}
		usr, err := hCtx.userStore.GetByEmail(cred.Email)
		if err == users.ErrUserNotFound { // maybe check if not nil
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		if err := usr.Authenticate(cred.Password); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		if _, err := sessions.BeginSession(hCtx.signingKey, hCtx.sessionStore, &SessionState{AuthUsr: usr, StartTime: time.Now()}, w); err != nil {
			http.Error(w, fmt.Sprintf("Error initiating session: %v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK, usr)
	} else {
		http.Error(w, "Only allowed to POST to this resource", http.StatusMethodNotAllowed)
		return
	}
}

func (hCtx *HandlerCtx) SessionsMineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		if _, err := sessions.EndSession(r, hCtx.signingKey, hCtx.sessionStore); err != nil {
			http.Error(w, fmt.Sprintf("Error getting session state: %v", err), http.StatusUnauthorized)
			return
		}
		w.Header().Add(headerContentType, contentTypePlainText)
		w.Write([]byte("signed out\n"))
		return
	} else {
		http.Error(w, "Only allowed to DELETE from this resource", http.StatusMethodNotAllowed)
		return
	}
}
