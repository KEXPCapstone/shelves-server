package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/KEXPCapstone/shelves-server/shelves/models"
	"gopkg.in/mgo.v2/bson"
)

func getUserIDFromRequest(r *http.Request) (bson.ObjectId, nil) {
	xUserHeader := r.Header.Get(XUser)
	if len(xUserHeader) == 0 || !bson.IsObjectIdHex(xUserHeader) {
		return nil, errors.New("NOT AUTHENTICATED")
	}
	userID := bson.ObjectIdHex(xUserHeader)
	return userID, nil
}

// /v1/shelves/mine/
func (hCtx *HandlerCtx) ShelvesMineHandler(w http.ResponseWriter, r *http.Request) {
	// used for getting just this specific user's shelves
}

// /v1/shelves
func (hCtx *HandlerCtx) ShelvesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		hCtx.getAllShelves(w, r)
	case http.MethodPost:
		hCtx.addShelf(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (hCtx *HandlerCtx) getAllShelves(w http.ResponseWriter, r *http.Request) {
	releases, err := hCtx.shelfStore.GetShelves()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error returned fetching shelves: %v", err), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusOK, releases)
}

func (hCtx *HandlerCtx) addShelf(w http.ResponseWriter, r *http.Request) {
	ns := &models.NewShelf{}
	userID, err := getUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "NOT AUTHENTICATED", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(ns); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON into new shelf: %v", err), http.StatusBadRequest)
		return
	}
	shelf, err := hCtx.shelfStore.InsertNew(ns, userID)
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
