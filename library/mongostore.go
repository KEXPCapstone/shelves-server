package library

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// implements LibraryStore interface
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

func (ms *MgoStore) GetReleaseByID(id bson.ObjectId) (*Release, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	release := &Release{}
	if err := coll.FindId(id).One(release); err != nil {
		return nil, ErrReleaseNotFound
	}
	return release, nil
}

func (ms *MgoStore) GetReleasesByKEXPCategory(KEXPCategory string) ([]*Release, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	releases := []*Release{}
	if err := coll.Find(bson.M{"KEXPCategory": KEXPCategory}).All(&releases); err != nil {
		return nil, ErrCategoryNotFound
	}
	return releases, nil

}
