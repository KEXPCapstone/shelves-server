package library

type LibraryStore interface {

	// TODO: return single track
	// Track will have associated release ID
	getTrackById(id int) error

	// TODO: return single release
	// TODO: decide if release has trackIds or track structs
	// Release will have associated artist ID
	getReleaseById(id int) error

	// TODO: Evaluate if necessary
	// Will return associated release
	getReleaseByTrackId(id int) error

	// TODO: Return a slice of releases within provided category
	getReleasesByCategory() error

	// TODO: Rename method?
	// TODO: Return a slice of releases which share the criterion with
	// the given releaseID
	// TODO: "SELECT * FROM Releases WHERE criterion == criterion" with
	// a switch on criterion
	getRelatedReleasesByReleaseId(id int, criterion string) error
}
