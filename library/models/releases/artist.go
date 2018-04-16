package releases

// Artist represents an artist in the KEXP library, with additional
// information about associated releases
type Artist struct {
	ID       string    `json:"id" bson:"_id"`
	Name     string    `json:"name" bson:"name"`
	Releases []Release `json:"releases" bson:"releases"`
}
