package library

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
