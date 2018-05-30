package models

import (
	"fmt"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// implements ShelvesStore interface
type MgoStore struct {
	session *mgo.Session
	dbname  string
	colname string
}

func NewMgoStore(sess *mgo.Session, dbName string, collectionName string) *MgoStore {
	if sess == nil {
		panic(NoMgoSess)
	}
	return &MgoStore{
		session: sess,
		dbname:  dbName,
		colname: collectionName,
	}
}

func (ms *MgoStore) InsertNew(ns *NewShelf, userId bson.ObjectId, ownerName string) (*Shelf, error) {
	shelf, err := ns.ToShelf(userId, ownerName)
	if err != nil {
		return nil, err
	}
	return ms.Insert(shelf)
}

func (ms *MgoStore) Insert(shelf *Shelf) (*Shelf, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	if err := coll.Insert(shelf); err != nil {
		return nil, fmt.Errorf("%v %v", ErrInsertShelf, err)
	}
	return shelf, nil
}

func (ms *MgoStore) GetShelves() ([]*Shelf, error) {
	shelves := []*Shelf{}
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	if err := coll.Find(nil).Sort("-datecreated").All(&shelves); err != nil {
		return nil, fmt.Errorf("%v %v", ErrShelfNotFound, err)
	}
	return shelves, nil
}

func (ms *MgoStore) GetShelfById(id bson.ObjectId) (*Shelf, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	shelf := &Shelf{}
	if err := coll.FindId(id).One(shelf); err != nil {
		return nil, fmt.Errorf("%v %v", ErrShelfNotFound, err)
	}
	return shelf, nil
}

func (ms *MgoStore) GetShelvesContainingRelease(id string) ([]*Shelf, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	results := []*Shelf{}
	if err := coll.Find(bson.M{"releases._id": id}).Select(bson.M{"_id": 1, "name": 1, "ownername": 1}).All(&results); err != nil {
		return nil, fmt.Errorf("%v %v", ErrShelfNotFound, err)
	}
	return results, nil
}

func (ms *MgoStore) GetUserShelves(userId bson.ObjectId) ([]*Shelf, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	shelves := []*Shelf{}
	if err := coll.Find(bson.M{"ownerid": userId}).Sort("-datecreated").All(&shelves); err != nil {
		return nil, fmt.Errorf("%v %v", ErrShelfNotFound, err)
	}
	return shelves, nil
}

func (ms *MgoStore) UpdateShelf(id bson.ObjectId, replacementShelf *Shelf) error {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	if err := coll.UpdateId(id, replacementShelf); err != nil {
		return err
	}
	return nil
}

func (ms *MgoStore) DeleteShelf(id bson.ObjectId) error {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	if err := coll.RemoveId(id); err != nil {
		return fmt.Errorf("%v %v", ErrDeleteShelf, err)
	}
	return nil
}

func (ms *MgoStore) CopyShelf(id bson.ObjectId, userId bson.ObjectId) (*Shelf, error) {
	shelf, err := ms.GetShelfById(id)
	if err != nil {
		return nil, err
	}
	shelf.ID = bson.NewObjectId()
	shelf.OwnerID = userId
	copied, err := ms.Insert(shelf)
	if err != nil {
		return nil, err
	}
	return copied, nil
}

func (ms *MgoStore) ExportShelf(id bson.ObjectId) error {
	return nil
}

func (ms *MgoStore) GetFeaturedShelves() ([]*Shelf, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	shelves := []*Shelf{}
	if err := coll.Find(bson.M{"featured": true}).Sort("-datecreated").All(&shelves); err != nil {
		return nil, fmt.Errorf("%v %v", ErrShelfNotFound, err)
	}
	return shelves, nil
}
