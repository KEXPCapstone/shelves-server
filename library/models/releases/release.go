package releases

import (
	"errors"
	"time"

	"github.com/KEXPCapstone/shelves-server/library/indexes"

	"github.com/globalsign/mgo/bson"
)

// Release represents a single release for an album in the KEXP library
// with a combination of fields from KEXP's own records, MetaData from MusicBrainz
// and custom fields related to the 'shelves' application
type Release struct {
	ID                      string          `json:"id" bson:"_id"`
	ArtistCredit            []interface{}   `json:"artistCredit" bson:"artist-credit"`
	ReleaseEvents           []interface{}   `json:"releaseEvents" bson:"release-events"`
	CoverArtArchive         interface{}     `json:"coverArtArchive" bson:"cover-art-archive"`
	KEXPReleaseGroupMBID    string          `json:"KEXPReleaseGroupMBID" bson:"KEXPReleaseGroupMBID"`
	KEXPReleaseArtistCredit string          `json:"KEXPReleaseArtistCredit" bson:"KEXPReleaseArtistCredit"`
	LabelInfo               []interface{}   `json:"labelInfo" bson:"label-info"`
	Media                   []interface{}   `json:"media" bson:"media"`
	Status                  string          `json:"status" bson:"status"`
	Disambiguation          string          `json:"disambiguation" bson:"disambiguation"`
	Barcode                 string          `json:"barcode" bson:"barcode"`
	Packaging               string          `json:"packaging" bson:"packaging"`
	Date                    string          `json:"date" bson:"date"`
	ASIN                    string          `json:"asin" bson:"asin"`
	KEXPPrimaryGenre        string          `json:"KEXPPrimaryGenre" bson:"KEXPPrimaryGenre"`
	Title                   string          `json:"title" bson:"title"`
	CountryCode             string          `json:"countryCode" bson:"country"`
	Yellows                 []int           `json:"yellows" bson:"yellows"`
	Reds                    []int           `json:"reds" bson:"reds"`
	Notes                   []bson.ObjectId `json:"noteIDs" bson:"noteIDs"`
}

// ReleaseAndMatchCriteria represents a pairing of
// a release object and a match criteria from a search index
type ReleaseAndMatchCriteria struct {
	Release   *Release             `json:"release"`
	IndexInfo indexes.SearchResult `json:"indexInfo"`
}

type NewNote struct {
	Comment string `json:"comment"`
}

// Note represents a note/comment for a given release
type Note struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	OwnerID     bson.ObjectId `json:"ownerID" bson:"ownerID"`
	AuthorName  string        `json:"authorName" bson:"authorName"`
	ReleaseID   string        `json:"releaseID" bson:"releaseID"`
	Comment     string        `json:"comment" bson:"comment"`
	DateCreated time.Time     `json:"dateCreated" bson:"dateCreated"`
	// DateLastEdit time.Time     `json:"dateLastEdit"`
}

func (nn *NewNote) Validate() error {
	if len(nn.Comment) == 0 {
		return errors.New(ErrEmptyComment)
	}
	return nil
}

func (nn *NewNote) ToNote(userID bson.ObjectId, authorName string, releaseID string) (*Note, error) {
	if err := nn.Validate(); err != nil {
		return nil, err
	}
	note := &Note{
		ID:          bson.NewObjectId(),
		OwnerID:     userID,
		AuthorName:  authorName,
		ReleaseID:   releaseID,
		Comment:     nn.Comment,
		DateCreated: time.Now(),
	}
	return note, nil
}
