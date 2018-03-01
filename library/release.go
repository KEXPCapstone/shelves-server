package library

import "gopkg.in/mgo.v2/bson"

type Release struct {
	ID              bson.ObjectId   `json:"id" bson:"_id"`
	Artist          string          `json:"artist"`
	Title           string          `json:"title"`
	ReleaseYear     string          `json:"releaseYear"`
	Label           string          `json:"label"`
	MusicBrainzInfo MusicBrainzMeta `json:"musicBrainzInfo"`
	DiscogsURL      string          `json:"discogsURL"`
	DaletID         string          `json:"daletID"`
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
