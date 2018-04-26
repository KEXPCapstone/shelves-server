package models

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Shelf struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	OwnerID      bson.ObjectId `json:"ownerId"` // TODO: May have to also add bson tag here
	Name         string        `json:"name"`
	ReleaseIDs   []string      `json:"releaseIDs"`
	Description  string        `json:"description"` // Maybe
	DateCreated  time.Time     `json:"dateCreated"`
	DateLastEdit time.Time     `json:"dateLastEdit"`
	Featured     bool          `json:"featured"` // Maybe
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

func (ns *NewShelf) ToShelf(userID bson.ObjectId) (*Shelf, error) {
	if err := ns.Validate(); err != nil {
		return nil, err
	}
	shelf := Shelf{
		ID:           bson.NewObjectId(),
		OwnerID:      userID,
		Name:         ns.Name,
		ReleaseIDs:   []bson.ObjectId{},
		Description:  ns.Description,
		DateCreated:  time.Now(),
		DateLastEdit: time.Now(),
		Featured:     ns.Featured,
	}
	return &shelf, nil
}
