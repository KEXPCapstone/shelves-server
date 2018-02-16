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
