package main

import (
	"encoding/json"
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
		// insert
		release := &Release{}
		if err := json.NewDecoder(r.Body).Decode(release); err != nil {
			// TODO
		}
		if err := hCtx.releaseStore.Insert(release); err != nil {
			// TODO
		}
		// respond
	case http.MethodGet:
		if len(field) != 0 && len(value) != 0 {
			releases, err := hCtx.releaseStore.GetReleasesByField(field, value)
			if err != nil {
				// TODO
			}
			// respond
		} else {
			releases, err := hCtx.releaseStore.GetAllReleases()
			if err != nil {
				// TODO
			}
			// respond
		}
	}
}

func (hCtx *HandlerCtx) SingleReleaseHandler(w http.ResponseWriter, r *http.Request) {
	releaseID := path.Base(r.URL.String())
	if !bson.IsObjectIdHex(releaseID) {
		// TODO: invalid
		return
	}
	switch r.Method {
	case http.MethodGet:
		release, err := hCtx.releaseStore.GetReleaseByID(releaseID)
		if err != nil {
			// TODO
		}
		// respond with release
	}
}
