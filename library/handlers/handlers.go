package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/KEXPCapstone/shelves-server/library/models/releases"
	"gopkg.in/mgo.v2/bson"
)

// ReleasesHandler path: /v1/library/releases
func (hCtx *HandlerCtx) ReleasesHandler(w http.ResponseWriter, r *http.Request) {
	field := r.URL.Query().Get("field")
	value := r.URL.Query().Get("value")
	searchTerm := r.URL.Query().Get("q")

	switch r.Method {
	case http.MethodPost:
		hCtx.insertRelease(w, r)
	case http.MethodGet:
		if len(field) != 0 && len(value) != 0 && len(searchTerm) == 0 {
			hCtx.findReleasesByField(w, r, field, value)
		} else if len(searchTerm) != 0 {
			hCtx.prefixSearch(w, r, searchTerm)
		} else { // What if we want to show results as user types? If searchTerm == 0, then all results are returned
			hCtx.getAllReleases(w, r)
		}
	default:
		http.Error(w, ReleasesHandlerInvalidMethod, http.StatusMethodNotAllowed)
		return
	}
}

// SingleReleaseHandler path: /v1/library/releases/
func (hCtx *HandlerCtx) SingleReleaseHandler(w http.ResponseWriter, r *http.Request) {
	releaseID := path.Base(r.URL.String())
	if !bson.IsObjectIdHex(releaseID) {
		http.Error(w, ErrInvalidReleaseID, http.StatusBadRequest)
		return
	}
	releaseIDBson := bson.ObjectIdHex(releaseID)
	switch r.Method {
	case http.MethodGet:
		release, err := hCtx.libraryStore.GetReleaseByID(releaseIDBson)
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

// TODO: Will probably remove
func (hCtx *HandlerCtx) insertRelease(w http.ResponseWriter, r *http.Request) {
	release := &releases.Release{}
	if err := json.NewDecoder(r.Body).Decode(release); err != nil {
		http.Error(w, fmt.Sprintf(ErrDecodingJSON+"%v", err), http.StatusBadRequest)
		return
	}
	if err := hCtx.libraryStore.AddRelease(release); err != nil {
		http.Error(w, fmt.Sprintf(ErrInsertRelease+"%v", err), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusCreated, release)
}

func (hCtx *HandlerCtx) findReleasesByField(w http.ResponseWriter, r *http.Request, field string, value string) {
	releases, err := hCtx.libraryStore.GetReleasesByField(field, value)
	if err != nil {
		http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusBadRequest)
		return
	}
	respond(w, http.StatusOK, releases)
}

func (hCtx *HandlerCtx) prefixSearch(w http.ResponseWriter, r *http.Request, searchTerm string) {
	searchTerm = strings.ToLower(searchTerm)
	searchResults := hCtx.releaseTrie.SearchReleases(searchTerm, maxSearchResults)
	foundReleases, err := hCtx.libraryStore.GetReleasesBySliceSearchResults(searchResults)
	if err != nil {
		http.Error(w, fmt.Sprintf(ErrorSearching+"%v", err), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusOK, foundReleases)
}

// Probably won't use this--if used should be paginated
func (hCtx *HandlerCtx) getAllReleases(w http.ResponseWriter, r *http.Request) {
	releases, err := hCtx.libraryStore.GetReleases()
	if err != nil {
		http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusBadRequest)
		return
	}
	respond(w, http.StatusOK, releases)
}
