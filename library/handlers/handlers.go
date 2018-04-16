package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
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
		lastID := r.URL.Query().Get("last_id")
		limit := r.URL.Query().Get("limit")

		intLimit, err := strconv.Atoi(limit)
		if err != nil {
			// for now this is a 500
			http.Error(w, fmt.Sprintf("Could not convert 'limit' param value '%v' to integer", limit), http.StatusInternalServerError)
		}
		hCtx.libraryStore.GetReleases(bson.ObjectIdHex(lastID), intLimit)

	default:
		http.Error(w, fmt.Sprintf(HandlerInvalidMethod, r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// RelatedReleasesHandler path: /v1/library/releases/related
// :param: field, the field key to match on
// :param: value, the target value of the match field
func (hCtx *HandlerCtx) RelatedReleasesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		field := r.URL.Query().Get("field")
		value := r.URL.Query().Get("value")
		if len(field) != 0 && len(value) != 0 {
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

// ArtistsHandler path: /v1/library/artists
func (hCtx *HandlerCtx) ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	return
}

// GenresHandler path: /v1/library/genres
func (hCtx *HandlerCtx) GenresHandler(w http.ResponseWriter, r *http.Request) {
	return
}

// TODO: Will probably remove
func (hCtx *HandlerCtx) insertRelease(w http.ResponseWriter, r *http.Request) {
	release := &releases.Release{}
	if err := json.NewDecoder(r.Body).Decode(release); err != nil {
		http.Error(w, fmt.Sprintf(ErrDecodingJSON+"%v", err), http.StatusBadRequest)
		return
	}
	result, err := hCtx.libraryStore.AddRelease(release)
	if err != nil {
		http.Error(w, fmt.Sprintf(ErrInsertRelease+"%v", err), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusCreated, result)
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
