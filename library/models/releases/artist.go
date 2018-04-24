package releases

import (
	"github.com/satori/go.uuid"
)

// Artist represents an artist in the KEXP library, with additional
// information about associated releases
type Artist struct {
	ArtistMBID uuid.UUID        `json:"artistMBID" bson:"_id"`
	Name       string           `json:"name" bson:"name"`
	Releases   []ReleaseSummary `json:"releases" bson:"releases"`
}

// ReleaseSummary represents a summary of a given release with minimal metadata
type ReleaseSummary struct {
	ReleaseMBID uuid.UUID `json:"releaseMBID" bson:"ReleaseMBID"`
	Title       string    `json:"title" bson:"title"`
	// ... more fields to be added as needed
}
