package library

import "gopkg.in/mgo.v2/bson"

type ReleaseStore interface {

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
	GetReleasesByCategory(category string) ([]*Release, error)

	// TODO: Rename method?
	// TODO: Return a slice of releases which share the criterion with
	// the given releaseID
	// TODO: "SELECT * FROM Releases WHERE criterion == criterion" with
	// a switch on criterion
	// "GET /v1/releases/related/{id}"
	GetRelatedReleasesByReleaseId(id bson.ObjectId, criterion string) ([]*Release, error)
}
