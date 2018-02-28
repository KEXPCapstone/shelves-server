package users

import mgo "gopkg.in/mgo.v2"

// implements UserStore interface
type MgoStore struct {
	session *mgo.Session
	dbname  string
	colname string
}

func NewMgoStore(sess *mgo.Session, dbName string, collectionName string) *MgoStore {
	if sess == nil {
		panic("nil pointer passed from session")
	}
	return &MgoStore{
		session: sess,
		dbname:  dbName,
		colname: collectionName,
	}
}

func (ms *MgoStore) Insert(nu *NewUser) (*User, error) {
	// TODO: Convert nu to an "intermediate user"
	// Place user into db, returning the associated id
	// Add the id field into the user?

	// Alternatively, convert nu to User without a
	// value for id, only pass in the relevant fields to be
	// added to a row in the database, and then write the
	// returned id to the User
	return nil, nil
}

func (ms *MgoStore) GetByID(id int) (*User, error) {
	return nil, nil
}

func (ms *MgoStore) GetByEmail(email string) (*User, error) {
	return nil, nil
}

func (ms *MgoStore) GetByUserName(username string) (*User, error) {
	return nil, nil
}

func (ms *MgoStore) Update(userID int, updates *Updates) error {
	return nil
}

func (ms *MgoStore) Delete(userID int) error {
	return nil
}
