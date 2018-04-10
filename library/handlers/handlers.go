package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"gopkg.in/mgo.v2/bson"
)

// TODO: Implement handler function
func (hCtx *HandlerCtx) ReleasesHandler(w http.ResponseWriter, r *http.Request) {
	field := r.URL.Query().Get("field")
	value := r.URL.Query().Get("value")

	switch r.Method {
	case http.MethodPost:
		release := &Release{}
		if err := json.NewDecoder(r.Body).Decode(release); err != nil {
			http.Error(w, fmt.Sprintf(ErrDecodingJSON+"%v", err), http.StatusBadRequest)
			return
		}
		if err := hCtx.releaseStore.Insert(release); err != nil {
			http.Error(w, fmt.Sprintf(ErrInsertRelease+"%v", err), http.StatusInternalServerError)
			return
		}
		// respond
	case http.MethodGet:
		if len(field) != 0 && len(value) != 0 {
			releases, err := hCtx.releaseStore.GetReleasesByField(field, value)
			if err != nil {
				http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusBadRequest)
				return
			}
			// respond
		} else {
			releases, err := hCtx.releaseStore.GetAllReleases()
			if err != nil {
				http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusBadRequest)
				return
			}
			// respond
		}
	}
}

func (hCtx *HandlerCtx) SingleReleaseHandler(w http.ResponseWriter, r *http.Request) {
	releaseID := path.Base(r.URL.String())
	if !bson.IsObjectIdHex(releaseID) {
		http.Error(w, ErrInvalidReleaseID, http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		release, err := hCtx.releaseStore.GetReleaseByID(releaseID)
		if err != nil {
			http.Error(w, fmt.Sprintf(ErrFetchingRelease+"%v", err), http.StatusBadRequest)
			return
		}
		// respond with release
	}
}
