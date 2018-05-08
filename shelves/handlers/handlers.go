package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/KEXPCapstone/shelves-server/gateway/models/users"
	"github.com/KEXPCapstone/shelves-server/shelves/models"
	"github.com/globalsign/mgo/bson"
)

// /v1/shelves/mine/
// used for getting just this specific user's shelves
func (hCtx *HandlerCtx) ShelvesMineHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userID, idErr := getUserIDFromRequest(r)
		if idErr != nil {
			http.Error(w, fmt.Sprintf("%v", idErr), http.StatusBadRequest)
			return
		}
		hCtx.getUsersShelvesFromID(w, r, userID)
	default:
		http.Error(w, ShelvesMineHandlerMethodNotAllowed, http.StatusMethodNotAllowed)
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
		http.Error(w, ShelvesHandlerMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
}

// /v1/shelves/users/{userID}
func (hCtx *HandlerCtx) UserShelvesHandler(w http.ResponseWriter, r *http.Request) {
	userID := path.Base(r.URL.String())
	if !bson.IsObjectIdHex(userID) {
		http.Error(w, InvalidUserID, http.StatusBadRequest)
		return
	}
	userIDBson := bson.ObjectIdHex(userID)
	switch r.Method {
	case http.MethodGet:
		shelves, err := hCtx.shelfStore.GetUserShelves(userIDBson)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK, shelves)
	default:
		http.Error(w, UserShelvesHandlerMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
}

// /v1/shelves/{id}
func (hCtx *HandlerCtx) ShelfHandler(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodPut:
		hCtx.updateShelf(w, r, shelf)
	case http.MethodDelete:
		hCtx.deleteShelf(w, r, shelf)
	default:
		http.Error(w, ShelfHandlerMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
}

// v1/shelves/featured
func (hCtx *HandlerCtx) FeaturedShelvesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		shelves, err := hCtx.shelfStore.GetFeaturedShelves()
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK, shelves)
	default:
		http.Error(w, FeaturedShelvesHandlerMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
}

/***
	HELPER METHODS
***/

func currUserIsShelfOwner(r *http.Request, shelf *models.Shelf) bool {
	userID, err := getUserIDFromRequest(r)
	if err != nil {
		return false
	}
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

func getUserIDFromRequest(r *http.Request) (bson.ObjectId, error) {
	xUserHeader := r.Header.Get(XUser)
	usr := &users.User{}
	if err := json.Unmarshal([]byte(xUserHeader), usr); err != nil {
		return bson.NewObjectId(), fmt.Errorf("%v : %v", ErrDecodingJSON, err)
	}
	if len(usr.ID) == 0 {
		return bson.NewObjectId(), ErrInvalidXUser
	}
	return usr.ID, nil
}

// Returns the full name of the current user.  If current user is not authenticated,
// returns an error
func getNameFromRequest(r *http.Request) (string, error) {
	xUserHeader := r.Header.Get(XUser)
	usr := &users.User{}
	if err := json.Unmarshal([]byte(xUserHeader), usr); err != nil {
		return "", fmt.Errorf("%v : %v", ErrDecodingJSON, err)
	}
	return usr.FirstName + " " + usr.LastName, nil
}

// Given a specific user ID, get that user's shelves.
func (hCtx *HandlerCtx) getUsersShelvesFromID(w http.ResponseWriter, r *http.Request, userID bson.ObjectId) {
	releases, err := hCtx.shelfStore.GetUserShelves(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusOK, releases)
}

func (hCtx *HandlerCtx) getAllShelves(w http.ResponseWriter, r *http.Request) {
	releases, err := hCtx.shelfStore.GetShelves()
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusOK, releases)
}

func (hCtx *HandlerCtx) addShelf(w http.ResponseWriter, r *http.Request) {
	ns := &models.NewShelf{}
	userID, err := getUserIDFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	ownerName, err := getNameFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(ns); err != nil {
		http.Error(w, fmt.Sprintf("%v : %v", ErrDecodingJSON, err), http.StatusBadRequest)
		return
	}
	shelf, err := hCtx.shelfStore.InsertNew(ns, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	respond(w, http.StatusCreated, shelf)
}

func (hCtx *HandlerCtx) updateShelf(w http.ResponseWriter, r *http.Request, shelf *models.Shelf) {
	if !currUserIsShelfOwner(r, shelf) {
		http.Error(w, ErrMustBeOwnerToEdit, http.StatusBadRequest)
		return
	}
	replacementShelf := &models.Shelf{}
	if err := json.NewDecoder(r.Body).Decode(replacementShelf); err != nil {
		http.Error(w, fmt.Sprintf("%v : %v", ErrDecodingJSON, err), http.StatusBadRequest)
		return
	}

	if err := hCtx.shelfStore.UpdateShelf(shelf.ID, replacementShelf); err != nil {
		http.Error(w, fmt.Sprintf("%v : %v", ErrUpdateShelf, err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(UpdatedShelfConf))
}

func (hCtx *HandlerCtx) deleteShelf(w http.ResponseWriter, r *http.Request, shelf *models.Shelf) {
	if !currUserIsShelfOwner(r, shelf) {
		http.Error(w, ErrMustBeOwnerToDelete, http.StatusBadRequest)
		return
	}
	if err := hCtx.shelfStore.DeleteShelf(shelf.ID); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(DeletedShelfConf))
}
