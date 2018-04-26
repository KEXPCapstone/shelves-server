package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// ReleasesHandler path: /v1/library/releases
// :param: last_id, id of the last release the client saw, for pagination
// :param: limit, the maximum number of releases to return
func (hCtx *HandlerCtx) ReleasesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		http.Error(w, fmt.Sprintf(HandlerInvalidMethod, r.Method), http.StatusMethodNotAllowed)
		return
	case http.MethodGet:
		lastID := r.URL.Query().Get("last_id")
		limit := r.URL.Query().Get("limit")

		intLimit, err := strconv.Atoi(limit)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not convert 'limit' param value '%v' to integer", limit), http.StatusInternalServerError)
			return
		}
		if _, err := uuid.Parse(lastID); err != nil {
			http.Error(w, fmt.Sprintf("Invalid value for 'last_id': '%v' (should be MBID)", lastID), http.StatusBadRequest)
			return
		}
		releases, err := hCtx.libraryStore.GetReleases(lastID, intLimit)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not fetch releases: %v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK, releases)

	default:
		http.Error(w, fmt.Sprintf(HandlerInvalidMethod, r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// SearchHandler path: /v1/library/search
// :param: q, the search query
func (hCtx *HandlerCtx) SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("q")

	switch r.Method {
	case http.MethodGet:
		if len(searchTerm) != 0 {
			hCtx.prefixSearch(w, r, searchTerm)
		}
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
		hCtx.findReleasesByField(w, r, field, value)
	default:
		http.Error(w, fmt.Sprintf(HandlerInvalidMethod, r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// SingleReleaseHandler path: /v1/library/releases/
func (hCtx *HandlerCtx) SingleReleaseHandler(w http.ResponseWriter, r *http.Request) {
	releaseID := path.Base(r.URL.String())
	if _, err := uuid.Parse(releaseID); err != nil {
		http.Error(w, fmt.Sprintf("'%v' is not a valid release id", releaseID), http.StatusBadRequest)
	}
	switch r.Method {
	case http.MethodGet:
		release, err := hCtx.libraryStore.GetReleaseByID(releaseID)
		if err != nil {
			http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK, release)
	default:
		http.Error(w, fmt.Sprintf(HandlerInvalidMethod, r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// ArtistsHandler path: /v1/library/artists
// :param: last_id, the id of the last artist (string artist name)
// :param: limit, the max number of artists to return
func (hCtx *HandlerCtx) ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		lastID := r.URL.Query().Get("last_id")
		limit := r.URL.Query().Get("limit")

		intLimit, err := strconv.Atoi(limit)
		if err != nil {
			// for now this is a 500
			http.Error(w, fmt.Sprintf("Could not convert 'limit' param value '%v' to integer", limit), http.StatusInternalServerError)
			return
		}
		releases, err := hCtx.libraryStore.GetArtists(lastID, intLimit)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not get artists: %v", err), http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK, releases)

	default:
		http.Error(w, fmt.Sprintf(HandlerInvalidMethod, r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// GenresHandler path: /v1/library/genres
func (hCtx *HandlerCtx) GenresHandler(w http.ResponseWriter, r *http.Request) {
	return
}

// /v1/library/notes/releases/{id}
func (hCtx *HandlerCtx) NotesHandler(w http.ResponseWriter, r *http.Request) {

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
