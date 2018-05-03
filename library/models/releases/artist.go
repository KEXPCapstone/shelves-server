package releases

// Artist represents an artist in the KEXP library, with additional
// information about associated releases
type Artist struct {
	ArtistMBID     string        `json:"artistMBID" bson:"_id"`
	ArtistName     string        `json:"name" bson:"artistName"`
	ArtistSortName string        `json:"artistSortName" bson:"artistSortName"`
	Disambiguation string        `json:"disambiguation" bson:"disambiguation"`
	ReleaseGroups  []interface{} `json:"releaseGroups" bson:"releaseGroups"`
}
