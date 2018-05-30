package models

import (
	"errors"
	"time"

	"github.com/KEXPCapstone/shelves-server/library/models/releases"

	"github.com/globalsign/mgo/bson"
)

type Shelf struct {
	ID           bson.ObjectId       `json:"id" bson:"_id"`
	OwnerID      bson.ObjectId       `json:"ownerId"`
	OwnerName    string              `json:"ownerName"`
	Name         string              `json:"name"`
	Releases     []*releases.Release `json:"releases"`
	Description  string              `json:"description"` // Maybe
	DateCreated  time.Time           `json:"dateCreated"`
	DateLastEdit time.Time           `json:"dateLastEdit"`
	Featured     bool                `json:"featured"`
}

type NewShelf struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Featured    bool   `json:"featured"`
}

func (ns *NewShelf) Validate() error {
	if len(ns.Name) == 0 {
		return errors.New(ErrEmptyShelfName)
	}
	return nil
}

func (ns *NewShelf) ToShelf(userID bson.ObjectId, ownerName string) (*Shelf, error) {
	if err := ns.Validate(); err != nil {
		return nil, err
	}
	shelf := Shelf{
		ID:           bson.NewObjectId(),
		OwnerID:      userID,
		OwnerName:    ownerName,
		Name:         ns.Name,
		Releases:     []*releases.Release{},
		Description:  ns.Description,
		DateCreated:  time.Now(),
		DateLastEdit: time.Now(),
		Featured:     ns.Featured,
	}
	return &shelf, nil
}
