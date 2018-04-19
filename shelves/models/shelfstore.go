package models

import "gopkg.in/mgo.v2/bson"

// TODO: Evaluate correct type of 'id' parameters
type ShelfStore interface {
	InsertNew(ns *NewShelf, userId bson.ObjectId) (*Shelf, error)

	Insert(shelf *Shelf) (*Shelf, error)

	// TODO: return array of shelves
	// "GET /v1/shelves/"
	GetShelves() ([]*Shelf, error)

	GetUserShelves(userId bson.ObjectId) ([]*Shelf, error)

	// TODO: return single shelf
	// "GET /v1/shelves/{id}"
	GetShelfById(id bson.ObjectId) (*Shelf, error)

	// TODO: Will accept a typeof 'Shelf' and replace exisiting shelf
	// at id with that new Shelf
	// "PUT /v1/shelves/{id}"
	UpdateShelf(id bson.ObjectId) error

	// "DELETE /v1/shelves/{id}"
	DeleteShelf(id bson.ObjectId) error

	// TODO: create and return copy of shelf with given id
	// TODO: Create resource path for this function
	CopyShelf(id bson.ObjectId, userId bson.ObjectId) (*Shelf, error)

	// TODO: export provided shelf as a folder to Dalet
	// This should probably be a handler function and not a store function
	ExportShelf(id bson.ObjectId) error
}
