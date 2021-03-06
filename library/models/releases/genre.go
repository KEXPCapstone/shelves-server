package releases

import (
	"github.com/globalsign/mgo/bson"
)

// Genre represents a genre in the KEXP library, with additional
// information about associated releases
type Genre struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Releases []Release     `json:"releases" bson:"releases"`
}
