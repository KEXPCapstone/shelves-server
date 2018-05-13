package models

import "github.com/globalsign/mgo/bson"

// TODO: Evaluate correct type of 'id' parameters
type ShelfStore interface {
	InsertNew(ns *NewShelf, userId bson.ObjectId, ownerName string) (*Shelf, error)

	Insert(shelf *Shelf) (*Shelf, error)

	// "GET /v1/shelves/"
	GetShelves() ([]*Shelf, error)

	GetUserShelves(userId bson.ObjectId) ([]*Shelf, error)

	// "GET /v1/shelves/{id}"
	GetShelfById(id bson.ObjectId) (*Shelf, error)

	// "PUT /v1/shelves/{id}"
	UpdateShelf(id bson.ObjectId, replacementShelf *Shelf) error

	// "DELETE /v1/shelves/{id}"
	DeleteShelf(id bson.ObjectId) error

	// TODO: Create resource path for this function
	CopyShelf(id bson.ObjectId, userId bson.ObjectId) (*Shelf, error)

	// TODO: export provided shelf as a folder to Dalet
	// This should probably be a handler function and not a store function
	ExportShelf(id bson.ObjectId) error

	// "GET /v1/shelves/featured"
	GetFeaturedShelves() ([]*Shelf, error)
}
