package main

import "gopkg.in/mgo.v2/bson"

type ReleaseStore interface {
	GetAllReleases() ([]*Release, error)

	// TODO: return single release
	// TODO: decide if release has trackIds or track structs
	// Release will have associated artist ID
	// "GET /v1/library/releases/{id}"
	GetReleaseByID(id bson.ObjectId) (*Release, error)

	// // TODO: Evaluate if necessary
	// // Will return associated release
	// // TODO: Resource path
	// getReleaseByTrackId(id int) error

	// TODO: Return a slice of releases within provided category
	// "GET /v1/library/releases/categories/{cat}"
	GetReleasesByField(field string, value string) ([]*Release, error)
}
