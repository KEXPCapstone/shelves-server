package releases

import (
	"time"

	"github.com/KEXPCapstone/shelves-server/library/indexes"
	"gopkg.in/mgo.v2/bson"
)

// type Release struct {
// 	ID              bson.ObjectId   `json:"id" bson:"_id"`
// 	Artist          string          `json:"artist"`
// 	Title           string          `json:"title"`
// 	ReleaseYear     string          `json:"releaseYear"`
// 	KEXPCategory    string          `json:"KEXPCategory"`
// 	Label           string          `json:"label"`
// 	MusicBrainzInfo MusicBrainzMeta `json:"musicBrainzInfo"`
// 	DiscogsURL      string          `json:"discogsURL"`
// 	DaletID         string          `json:"daletID"`
// 	TrackList       []Track         `json:"trackList"`
// 	Notes           []Note          `json:"notes"`
// 	// Add image URL?
// }

// Changed name from DaletRelease for testing
type Release struct {
	ID                            bson.ObjectId `json:"id" bson:"_id"`
	KEXPPrimaryGenre              string        `json:"KEXPPrimaryGenre" bson:"KEXPPrimaryGenre"`
	KEXPMBID                      string        `json:"KEXPMBID" bson:"KEXPMBID"`
	KEXPDateReleased              string        `json:"KEXPDateReleased" bson:"KEXPDateReleased"`
	KEXPFirstReleaseDate          string        `json:"KEXPFirstReleaseDate" bson:"KEXPFirstReleaseDate"`
	KEXPLength                    string        `json:"KEXPLength" bson:"KEXPLength"`
	KEXPReleaseCatalogNumber      string        `json:"KEXPReleaseCatalogNumber" bson:"KEXPReleaseCatalogNumber"` // couldn't find
	KEXPReleaseGroupMBID          string        `json:"KEXPReleaseGroupMBID" bson:"KEXPReleaseGroupMBID"`
	KEXPTitle                     string        `json:"KEXPTitle" bson:"KEXPTitle"`
	KEXPUniqueTitle               string        `json:"KEXPUniqueTitle" bson:"KEXPUniqueTitle"`
	KEXPReleaseArtistCredit       string        `json:"KEXPReleaseArtistCredit" bson:"KEXPReleaseArtistCredit"`
	KEXPArtist                    string        `json:"KEXPArtist" bson:"KEXPArtist"`
	KEXPLabel                     string        `json:"KEXPLabel" bson:"KEXPLabel"`
	KEXPReleasePackaging          string        `json:"KEXPReleasePackaging" bson:"KEXPReleasePackaging"`
	KEXPReleasePrimaryType        string        `json:"KEXPReleasePrimaryType" bson:"KEXPReleasePrimaryType"`
	KEXPReleaseSecondaryType      string        `json:"KEXPReleaseSecondaryType" bson:"KEXPReleaseSecondaryType"`
	KEXPReleaseStatus             string        `json:"KEXPReleaseStatus" bson:"KEXPReleaseStatus"`
	KEXPArtist_KEXPAlias          string        `json:"KEXPArtist_KEXPAlias" bson:"KEXPArtist_KEXPAlias"`
	KEXPArtist_KEXPArtistType     string        `json:"KEXPArtist_KEXPArtistType" bson:"KEXPArtist_KEXPArtistType"`
	KEXPArtist_KEXPDisambiguation string        `json:"KEXPArtist_KEXPDisambiguation" bson:"KEXPArtist_KEXPDisambiguation"`
	KEXPArtist_KEXPLink           string        `json:"KEXPArtist_KEXPLink" bson:"KEXPArtist_KEXPLink"`
	KEXPArtist_KEXPMBID           string        `json:"KEXPArtist_KEXPMBID" bson:"KEXPArtist_KEXPMBID"`
	KEXPArtist_KEXPName           string        `json:"KEXPArtist_KEXPName" bson:"KEXPArtist_KEXPName"`
	KEXPArtist_KEXPSortName       string        `json:"KEXPArtist_KEXPSortName" bson:"KEXPArtist_KEXPSortName"`
	KEXPLabel_KEXPMBID            string        `json:"KEXPLabel_KEXPMBID" bson:"KEXPLabel_KEXPMBID"`
	KEXPLabel_KEXPName            string        `json:"KEXPLabel_KEXPName" bson:"KEXPLabel_KEXPName"`
	KEXPLabel_KEXPSortName        string        `json:"KEXPLabel_KEXPSortName" bson:"KEXPLabel_KEXPSortName"`
	KEXPArea                      string        `json:"KEXPArea" bson:"KEXPArea"`
	KEXPAreaMBID                  string        `json:"KEXPAreaMBID" bson:"KEXPAreaMBID"`
	KEXPCountryCode               string        `json:"KEXPCountryCode" bson:"KEXPCountryCode"`
}

type MusicBrainzMeta struct {
	MBID string `json:"mbid"`
	// TODO: All of the meta that we want.
}

type Track struct {
	Name      string `json:"name"`
	Length    string `json:"length"`
	FCCRating string `json:"FCCRating"`
}

type ReleaseAndMatchCriteria struct {
	Release   *Release             `json:"release"`
	IndexInfo indexes.SearchResult `json:"indexInfo"`
}

type Note struct {
	Author      string        `json:"author"`
	Comment     string        `json:"comment"`
	UserID      bson.ObjectId `json:"userID"`
	DateCreated time.Time     `json:"dateCreated"`
	// TODO: Decide if editing notes is out of scope or not
	// DateLastEdit time.Time     `json:"dateLastEdit"`
}

// func (dr *DaletRelease) ProcessDaletRelease() (*Release, error) {

// 		1). Load in Data from Dalet into DaletRelease
// 			a). Handle scenarios of missing data.
// 		2). Fetch data from MusicBrainz
// 			a). Producers
// 			b). Building track lists?
// 			c). Any other fields
// 		3). Copy fields into Release struct with associated BSON ID
// 		4). Place Release into Mongo.

// 	return nil, nil
// }
