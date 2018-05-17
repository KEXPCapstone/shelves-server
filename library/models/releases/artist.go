package releases

// Artist represents an artist in the KEXP library, with additional
// information about associated releases
type Artist struct {
	ArtistMBID     string        `json:"id" bson:"_id"`
	ArtistName     string        `json:"name" bson:"artistName"`
	ArtistSortName string        `json:"sortName" bson:"artistSortName"`
	Disambiguation string        `json:"disambiguation" bson:"disambiguation"`
	ReleaseGroups  []interface{} `json:"releaseGroups" bson:"releaseGroups"`
}
