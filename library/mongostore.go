package library

import mgo "gopkg.in/mgo.v2"

// implements LibraryStore interface
type MgoStore struct {
	session *mgo.Session
	dbname  string
	colname string
}
