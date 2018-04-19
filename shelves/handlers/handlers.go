package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.com/KEXPCapstone/shelves-server/shelves/models"
	"gopkg.in/mgo.v2/bson"
)

// /v1/shelves/mine/
// used for getting just this specific user's shelves
func (hCtx *HandlerCtx) ShelvesMineHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userID, idErr := getUserIDFromRequest(r)
		if idErr != nil {
			http.Error(w, "NOT AUTHENTICATED", http.StatusBadRequest)
			return
		}
		hCtx.getUsersShelvesFromID(w, r, userID)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
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

// /v1/shelves/{id}
func (hCtx *HandlerCtx) ShelfHandler(w http.ResponseWriter, r *http.Request) {
	// Refactor?
	shelf, err := hCtx.getShelfFromRequest(r)
	if err == ErrInvalidShelfID {
		http.Error(w, fmt.Sprintf("%v", ErrInvalidShelfID), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, http.StatusOK, shelf)
	case http.MethodPatch:
		hCtx.updateShelf(w, r, shelf)
	case http.MethodDelete:
		hCtx.deleteShelf(w, r, shelf)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

/***
	HELPER METHODS
***/

func currUserIsShelfOwner(r *http.Request, shelf *models.Shelf) boolean {
	userID, err := getUserIDFromRequest(r)
	if userID != shelf.OwnerID {
		return false
	}
	return true
}

func (hCtx *HandlerCtx) getShelfFromRequest(r *http.Request) (*models.Shelf, error) {
	shelfID := path.Base(r.URL.String())
	if !bson.IsObjectIdHex(shelfID) {
		return nil, ErrInvalidShelfID
	}
	shelfIDBson := bson.ObjectIdHex(shelfID)
	shelf, err := hCtx.shelfStore.GetShelfById(shelfIDBson)
	if err != nil {
		return nil, err
	}
	return shelf, nil
}

func getUserIDFromRequest(r *http.Request) (bson.ObjectId, nil) {
	xUserHeader := r.Header.Get(XUser)
	if len(xUserHeader) == 0 || !bson.IsObjectIdHex(xUserHeader) {
		return nil, errors.New("NOT AUTHENTICATED")
	}
	userID := bson.ObjectIdHex(xUserHeader)
	return userID, nil
}

func (hCtx *HandlerCtx) getUsersShelvesFromID(w http.ResponseWriter, r *http.Request, userID bson.ObjectId) {
	releases, err := hCtx.shelfStore.GetUserShelves(userID)
	if err != nil {
		http.Error(w, fmt.Sprint("Error returned fetching user's shelves: %v", err), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusOK, releases)
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

func (hCtx *HandlerCtx) updateShelf(w http.ResponseWriter, r *http.Request, shelf *models.Shelf) {
	if !currUserIsShelfOwner(r, shelf) {
		http.Error(w, "You must own the shelf to edit it", http.StatusBadRequest)
		return
	}
	if err := hCtx.shelfStore.UpdateShelf(shelfID); err != nil {
		// Undefined behavior
	}
	// respond
}

func (hCtx *HandlerCtx) deleteShelf(w http.ResponseWriter, r *http.Request, shelf *models.Shelf) {
	if !currUserIsShelfOwner(r, shelf) {
		http.Error(w, "You must own the shelf to delete it", http.StatusBadRequest)
		return
	}
	if err := hCtx.shelfStore.DeleteShelf(shelfID); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Deleted shelf\n"))
}
