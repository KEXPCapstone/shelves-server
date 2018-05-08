package users

import (
	"fmt"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// implements UserStore interface
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

func (ms *MgoStore) Insert(nu *NewUser) (*User, error) {
	usr, err := nu.ToUser()
	if err != nil {
		return nil, err
	}
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	if err := coll.Insert(usr); err != nil {
		return nil, ErrInsertUser
	}
	return usr, nil
}

func (ms *MgoStore) GetByID(id bson.ObjectId) (*User, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	usr := User{}
	if err := coll.FindId(id).One(&usr); err != nil {
		return nil, ErrUserNotFound
	}
	return &usr, nil

}

func (ms *MgoStore) GetByEmail(email string) (*User, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	usr := User{}
	if err := coll.Find(bson.M{"email": email}).One(&usr); err != nil {
		return nil, ErrUserNotFound
	}
	return &usr, nil
}

func (ms *MgoStore) Delete(id bson.ObjectId) error {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	if err := coll.RemoveId(id); err != nil {
		return fmt.Errorf("%v : %v", ErrStrDeleteUser, err)
	}
	return nil
}
