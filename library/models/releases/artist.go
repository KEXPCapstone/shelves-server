package releases

// Artist represents an artist in the KEXP library, with additional
// information about associated releases
type Artist struct {
	ID       string           `json:"id" bson:"_id"`
	Name     string           `json:"name" bson:"name"`
	Releases []ReleaseSummary `json:"releases" bson:"releases"`
}

// ReleaseSummary represents a summary of a given release with minimal metadata
type ReleaseSummary struct {
	KEXPMBID  string `json:"KEXPMBID" bson:"KEXPMBID"`
	KEXPTitle string `json:"KEXPTitle" bson:"KEXPTitle"`
	// ... more fields to be added as needed
}
