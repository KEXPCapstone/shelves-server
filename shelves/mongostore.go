package shelves

import mgo "gopkg.in/mgo.v2"

// implements ShelvesStore interface
type MgoStore struct {
	session *mgo.Session
	dbname  string
	colname string
}
