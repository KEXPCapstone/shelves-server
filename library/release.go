package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Release struct {
	ID              bson.ObjectId   `json:"id" bson:"_id"`
	Artist          string          `json:"artist"`
	Title           string          `json:"title"`
	ReleaseYear     string          `json:"releaseYear"`
	KEXPCategory    string          `json:"KEXPCategory"`
	Label           string          `json:"label"`
	MusicBrainzInfo MusicBrainzMeta `json:"musicBrainzInfo"`
	DiscogsURL      string          `json:"discogsURL"`
	DaletID         string          `json:"daletID"`
	TrackList       []Track         `json:"trackList"`
	Notes           []Note          `json:"notes"`
	// Add image URL?
}

type DaletRelease struct {
	KEXPPrimaryGenre              string `json:"KEXPPrimaryGenre"`
	KEXPMBID                      string `json:"KEXPMBID"`
	KEXPDateReleased              string `json:"KEXPDateReleased"`
	KEXPFirstReleaseDate          string `json:"KEXPFirstReleaseDate"`
	KEXPLength                    string `json:"KEXPLength"`
	KEXPReleaseCatalogNumber      string `json:"KEXPReleaseCatalogNumber"` // couldn't find
	KEXPReleaseGroupMBID          string `json:"KEXPReleaseGroupMBID"`
	KEXPTitle                     string `json:"KEXPTitle"`
	KEXPUniqueTitle               string `json:"KEXPUniqueTitle"`
	KEXPReleaseArtistCredit       string `json:"KEXPReleaseArtistCredit"`
	KEXPArtist                    string `json:"KEXPArtist"`
	KEXPLabel                     string `json:"KEXPLabel"`
	KEXPReleasePackaging          string `json:"KEXPReleasePackaging"`
	KEXPReleasePrimaryType        string `json:"KEXPReleasePrimaryType"`
	KEXPReleaseSecondaryType      string `json:"KEXPReleaseSecondaryType"`
	KEXPReleaseStatus             string `json:"KEXPReleaseStatus"`
	KEXPArtist_KEXPAlias          string `json:"KEXPArtist_KEXPAlias"`
	KEXPArtist_KEXPArtistType     string `json:"KEXPArtist_KEXPArtistType"`
	KEXPArtist_KEXPDisambiguation string `json:"KEXPArtist_KEXPDisambiguation"`
	KEXPArtist_KEXPLink           string `json:"KEXPArtist_KEXPLink"`
	KEXPArtist_KEXPMBID           string `json:"KEXPArtist_KEXPMBID"`
	KEXPArtist_KEXPName           string `json:"KEXPArtist_KEXPName"`
	KEXPArtist_KEXPSortName       string `json:"KEXPArtist_KEXPSortName"`
	KEXPLabel_KEXPMBID            string `json:"KEXPLabel_KEXPMBID"`
	KEXPLabel_KEXPName            string `json:"KEXPLabel_KEXPName"`
	KEXPLabel_KEXPSortName        string `json:"KEXPLabel_KEXPSortName"`
	KEXPArea                      string `json:"KEXPArea"`
	KEXPAreaMBID                  string `json:"KEXPAreaMBID"`
	KEXPCountryCode               string `json:"KEXPCountryCode"`
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

type Note struct {
	Author      string        `json:"author"`
	Comment     string        `json:"comment"`
	UserID      bson.ObjectId `json:"userID"`
	DateCreated time.Time     `json:"dateCreated"`
	// TODO: Decide if editing notes is out of scope or not
	// DateLastEdit time.Time     `json:"dateLastEdit"`
}
