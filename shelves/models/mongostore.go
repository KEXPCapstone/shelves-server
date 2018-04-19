package models

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// implements ShelvesStore interface
type MgoStore struct {
	session *mgo.Session
	dbname  string
	colname string
	// TODO: Maybe just have a field for coll here
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

func (ms *MgoStore) InsertNew(ns *NewShelf, userId bson.ObjectId) (*Shelf, error) {
	shelf, err := ns.ToShelf(userId)
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
	if err := coll.Find(nil).All(&shelves); err != nil {
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

func (ms *MgoStore) GetUserShelves(userId bson.ObjectId) ([]*Shelf, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	shelves := []*Shelf{}
	if err := coll.Find(bson.M{"OwnerID": userId}).All(&shelves); err != nil {
		return nil, fmt.Errorf("%v %v", ErrShelfNotFound, err)
	}
	return shelves, nil
}

// TODO: Evaluate best means of updating; replacing or patching
func (ms *MgoStore) UpdateShelf(id bson.ObjectId) error {
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
