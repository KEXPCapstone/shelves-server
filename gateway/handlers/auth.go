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
	sid, err := sessions.GetState(r, hCtx.signingKey, hCtx.sessionStore, ss)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching current user: %v", err), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, http.StatusOK, ss.AuthUsr)
	case http.MethodPatch:
		upd := &users.Updates{}
		if err := json.NewDecoder(r.Body).Decode(upd); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON into Updates: %v", err), http.StatusBadRequest)
			return
		}
		if err := hCtx.userStore.Update(ss.AuthUsr.ID, upd); err != nil {
			http.Error(w, fmt.Sprintf("Error updating user: %v", err), http.StatusBadRequest)
			return
		}
		usr, err := hCtx.userStore.GetByID(ss.AuthUsr.ID) // get the newly updated user
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching updated usr: %v", err), http.StatusInternalServerError)
			return
		}
		if err := hCtx.sessionStore.Save(sid, &SessionState{AuthUsr: usr}); err != nil { // update sessions store
			http.Error(w, fmt.Sprintf("Error updating session: %v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK, usr)
	default:
		http.Error(w, "Only allowed to GET or PATCH to this resource", http.StatusMethodNotAllowed)
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
