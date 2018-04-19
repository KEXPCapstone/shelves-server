package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KEXPCapstone/shelves-server/shelves/models"
	"gopkg.in/mgo.v2/bson"
)

// /v1/shelves/mine/
func (hCtx *HandlerCtx) ShelvesMineHandler(w http.ResponseWriter, r *http.Request) {
	// used for getting just this specific user's shelves
}

// /v1/shelves
func (hCtx *HandlerCtx) ShelvesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		hCtx.addShelf(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (hCtx *HandlerCtx) addShelf(w http.ResponseWriter, r *http.Request) {
	ns := &models.NewShelf{}
	xUserHeader := r.Header.Get(XUser)
	if len(xUserHeader) == 0 || !bson.IsObjectIdHex(xUserHeader) {
		http.Error(w, "NOT AUTHENTICATED", http.StatusBadRequest)
		return
	}
	userId := bson.ObjectIdHex(xUserHeader)
	if err := json.NewDecoder(r.Body).Decode(ns); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON into new shelf: %v", err), http.StatusBadRequest)
		return
	}
	shelf, err := hCtx.shelfStore.InsertNew(ns, userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when adding new shelf: %v", err), http.StatusBadRequest)
		return
	}
	respond(w, http.StatusCreated, shelf)
}

// /v1/shelves/{id}
func (hCtx *HandlerCtx) ShelfHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}
