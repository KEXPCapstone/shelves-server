package shelves

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

func (ms *MgoStore) Insert(ns *NewShelf) (*Shelf, error) {
	userId := bson.NewObjectId() // TODO: will need to get userId from sessionState
	shelf, err := ns.ToShelf(userId)
	if err != nil {
		return nil, err
	}
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
		return nil, ErrFindShelf
	}
	return shelves, nil
}
