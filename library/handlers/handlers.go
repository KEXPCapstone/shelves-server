package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/KEXPCapstone/shelves-server/library/models/releases"
	"gopkg.in/mgo.v2/bson"
)

func (hCtx *HandlerCtx) ReleasesHandler(w http.ResponseWriter, r *http.Request) {
	field := r.URL.Query().Get("field")
	value := r.URL.Query().Get("value")
	searchTerm := r.URL.Query().Get("q")

	switch r.Method {
	case http.MethodPost:
		release := &releases.Release{}
		if err := json.NewDecoder(r.Body).Decode(release); err != nil {
			http.Error(w, fmt.Sprintf(ErrDecodingJSON+"%v", err), http.StatusBadRequest)
			return
		}
		if err := hCtx.releaseStore.Insert(release); err != nil {
			http.Error(w, fmt.Sprintf(ErrInsertRelease+"%v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusCreated, release)
	case http.MethodGet:
		if len(field) != 0 && len(value) != 0 && len(searchTerm) == 0 {
			log.Println(field)
			log.Println(value)
			releases, err := hCtx.releaseStore.GetReleasesByField(field, value)
			if err != nil {
				http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusBadRequest)
				return
			}
			respond(w, http.StatusOK, releases)
		} else if len(searchTerm) != 0 {
			searchTerm = strings.ToLower(searchTerm)
			foundReleaseIDs := hCtx.releaseTrie.SearchReleases(searchTerm, maxSearchResults)
			foundReleases, err := hCtx.releaseStore.GetReleasesBySliceObjectIDs(foundReleaseIDs)
			if err != nil {
				http.Error(w, fmt.Sprintf(ErrorSearching+"%v", err), http.StatusInternalServerError)
				return
			}
			respond(w, http.StatusOK, foundReleases)

		} else { // What if we want to show results as user types? If searchTerm == 0, then all results are returned
			releases, err := hCtx.releaseStore.GetAllReleases()
			if err != nil {
				http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusBadRequest)
				return
			}
			respond(w, http.StatusOK, releases)
		}
	default:
		http.Error(w, ReleasesHandlerInvalidMethod, http.StatusMethodNotAllowed)
		return
	}
}

func (hCtx *HandlerCtx) SingleReleaseHandler(w http.ResponseWriter, r *http.Request) {
	releaseID := path.Base(r.URL.String())
	if !bson.IsObjectIdHex(releaseID) {
		http.Error(w, ErrInvalidReleaseID, http.StatusBadRequest)
		return
	}
	releaseIDBson := bson.ObjectIdHex(releaseID)
	switch r.Method {
	case http.MethodGet:
		release, err := hCtx.releaseStore.GetReleaseByID(releaseIDBson)
		if err != nil {
			http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusBadRequest)
			return
		}
		respond(w, http.StatusOK, release)
	default:
		http.Error(w, SingleReleaseHandlerInvalidMethod, http.StatusMethodNotAllowed)
		return
	}
}
