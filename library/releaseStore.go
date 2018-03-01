package library

type ReleaseStore interface {

	// TODO: return single track
	// Track will have associated release ID
	// "GET /v1/library/tracks/{id}"
	getTrackById(id int) error

	// TODO: return single release
	// TODO: decide if release has trackIds or track structs
	// Release will have associated artist ID
	// "GET /v1/library/releases/{id}"
	getReleaseById(id int) error

	// TODO: Evaluate if necessary
	// Will return associated release
	// TODO: Resource path
	getReleaseByTrackId(id int) error

	// TODO: Return a slice of releases within provided category
	// "GET /v1/library/releases/categories/{cat}"
	getReleasesByCategory() error

	// TODO: Rename method?
	// TODO: Return a slice of releases which share the criterion with
	// the given releaseID
	// TODO: "SELECT * FROM Releases WHERE criterion == criterion" with
	// a switch on criterion
	// "GET /v1/releases/related/{id}"
	getRelatedReleasesByReleaseId(id int, criterion string) error
}
