package releases

import (
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

// a helper function which will create a new MongoDB database
// and collection.  Returns a pointer to the associated MgoStore or an error
func createTestingMgoStore() (*MgoStore, error) {
	mongoSess, err := mgo.Dial("localhost")
	if err != nil {
		return nil, fmt.Errorf("Error dialing MongoDB. Error: %v", err)
	}
	ms := NewMgoStore(mongoSess, "library", "releases")
	return ms, nil
}

func TestNewMgoStore(t *testing.T) {
	mongoSess, err := mgo.Dial("localhost")
	dbname := "library"
	colname := "releases"
	if err != nil {
		t.Errorf("Error dialing mongodb. Error: %v", err)
	}
	mgoStore := NewMgoStore(mongoSess, dbname, colname)
	if mgoStore == nil {
		t.Errorf("Mgostore returned nil")
	}
	if mgoStore.session != mongoSess {
		t.Errorf("Error setting session. Expected: '%v' Actual: '%v'", mongoSess, mgoStore.session)
	}
	if mgoStore.dbname != dbname {
		t.Errorf("Error setting database name.  Expected: '%v' Actual: '%v'", dbname, mgoStore.dbname)
	}
	if mgoStore.colname != colname {
		t.Errorf("Error setting collection name. Expected: '%v' Actual: '%v'", colname, mgoStore.colname)
	}
}

func TestGetReleaseByID(t *testing.T) {

}
