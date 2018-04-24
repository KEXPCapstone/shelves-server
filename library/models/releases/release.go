package releases

import (
	"time"

	"github.com/KEXPCapstone/shelves-server/library/indexes"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Release represents a single release for an album in the KEXP library
// with a combination of fields from KEXP's own records, MetaData from MusicBrainz
// and custom fields related to the 'shelves' application
type Release struct {
	ID                      uuid.UUID     `json:"id" bson:"_id"`
	ArtistCredit            []interface{} `json:"artistCredit" bson:"artist-credit"`
	ReleaseEvents           []interface{} `json:"releaseEvents" bson:"release-events"`
	CoverArtArchive         interface{}   `json:"coverArtArchive" bson:"cover-art-archive"`
	KEXPReleaseGroupMBID    string        `json:"KEXPReleaseGroupMBID" bson:"KEXPReleaseGroupMBID"`
	KEXPReleaseArtistCredit string        `json:"KEXPReleaseArtistCredit" bson:"KEXPReleaseArtistCredit"`
	LabelInfo               []interface{} `json:"labelInfo" bson:"label-info"`
	Media                   string        `json:"media" bson:"media"`
	Status                  string        `json: "status" bson: "status"`
	Disambiguation          string        `json: "disambiguation" bson: "disambiguation"`
	Barcode                 string        `json:"barcode" bson:"barcode"`
	Packaging               string        `json:"packaging" bson:"packaging"`
	Date                    string        `json:"date" bson:"date"`
	ASIN                    string        `json:"asin" bson:"asin"`
	KEXPPrimaryGenre        string        `json:"KEXPPrimaryGenre" bson:"KEXPPrimaryGenre"`
	Title                   string        `json:"title" bson:"title"`
	CountryCode             string        `json:"countryCode" bson:"country"`
	Yellows                 []int         `json:"yellows" bson:"yellows"`
	Reds                    []int         `json:"reds" bson:"reds"`
}

// ReleaseAndMatchCriteria represents a pairing of
// a release object and a match criteria from a search index
type ReleaseAndMatchCriteria struct {
	Release   *Release             `json:"release"`
	IndexInfo indexes.SearchResult `json:"indexInfo"`
}

// Note represents a note/comment for a given release
type Note struct {
	Author      string        `json:"author"`
	Comment     string        `json:"comment"`
	UserID      bson.ObjectId `json:"userID"`
	DateCreated time.Time     `json:"dateCreated"`
	// DateLastEdit time.Time     `json:"dateLastEdit"`
}
