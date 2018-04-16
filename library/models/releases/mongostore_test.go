package releases

import (
	"encoding/json"
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// a helper function which will create a new MongoDB database
// and collection.  Returns a pointer to the associated MgoStore or an error
func createTestingMgoStore() (*MongoStore, error) {
	mongoSess, err := mgo.Dial("localhost")
	if err != nil {
		return nil, fmt.Errorf("Error dialing MongoDB. Error: %v", err)
	}
	dbname := "library"
	releaseCol := "releases"
	artistCol := "artists"
	genreCol := "genres"
	ms := NewMongoStore(mongoSess, dbname, releaseCol, artistCol, genreCol)
	return ms, nil
}

func TestNewMongoStore(t *testing.T) {
	dbname := "library"
	releaseCol := "releases"
	artistCol := "artists"
	genreCol := "genres"
	mongoSess, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error dialing mongodb: %v", err)
	}
	libraryStore := NewMongoStore(mongoSess, dbname, releaseCol, artistCol, genreCol)
	if libraryStore.session != mongoSess {
		t.Errorf("Error setting session. Expected: '%v', Actual: '%v'", mongoSess, libraryStore.session)
	}
	if libraryStore.dbname != dbname {
		t.Errorf("Error setting database name.  Expected: '%v', Actual: '%v'", dbname, libraryStore.dbname)
	}
	if libraryStore.releaseCollection != releaseCol {
		t.Errorf("Error setting release collection name. Expected: '%v' Actual: '%v'", releaseCol, libraryStore.releaseCollection)
	}
	if libraryStore.artistCollection != artistCol {
		t.Errorf("Error setting artist collection name. Expected: '%v' Actual: '%v'", artistCol, libraryStore.artistCollection)
	}
	if libraryStore.genreCollection != genreCol {
		t.Errorf("Error setting genre collection name. Expected: '%v' Actual: '%v'", genreCol, libraryStore.genreCollection)
	}
}

func TestNewMongoStoreNilSession(t *testing.T) {
	// defer func to test panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// this call should cause panic() to be caught by the defer above
	_ = NewMongoStore(nil, "foo", "bar", "baz", "bip")
}

func TestAddRelease(t *testing.T) {
	libraryStore, err := createTestingMgoStore()
	if err != nil {
		t.Errorf("[MetaTest] Error creating test MongoStore: '%v", err)
	}
	newRelease, err := createDefaultTestRelease(bson.NewObjectId())
	if err != nil {
		t.Errorf("[MetaTest] %v", err)
	}
	if _, err = libraryStore.AddRelease(newRelease); err != nil {
		t.Errorf("Error inserting doc into 'releases' collection: %v", err)
	}
	if _, err = libraryStore.AddRelease(nil); err == nil {
		t.Errorf("Expected error inserting `nil` release doc")
	}
}

// helper function creates a new release struct populated with placeholder
// fields to be reset to whatever is needed
func createDefaultTestRelease(id bson.ObjectId) (*Release, error) {
	// unmarshal an example doc into release struct
	newRelease := &Release{
		ID: id,
	}
	// make this a constant
	rawJSON := `{"KEXPReleaseGroupMBID" : "62917949-8997-409b-94fb-af8a41ff3c52", "KEXPPrimaryGenre" : "Rock/Pop", "KEXPCountryCode" : "US", "KEXPDateReleased" : "2015-09-25", "KEXPMBID" : "cd7d006c-a9fa-4094-a733-6d001dcfa4b4", "KEXPReleaseArtistCredit" : "Kurt Vile", "KEXPArtist_KEXPSortName" : "Vile, Kurt", "KEXPFirstReleaseDate" : "2015-09-25", "KEXPTitle" : "b'lieve i'm goin down..."}`
	buffer := []byte(rawJSON)
	if err := json.Unmarshal(buffer, newRelease); err != nil {
		return nil, fmt.Errorf("[MetaTest] Error unmarshalling JSON: '%v'", err)
	}
	return newRelease, nil
}

func TestGetReleases(t *testing.T) {

}

func TestGetReleaseByID(t *testing.T) {
	libraryStore, err := createTestingMgoStore()
	if err != nil {
		t.Errorf("[MetaTest] Error creating test MongoStore: '%v", err)
	}
	// insert a new doc
	targetID := bson.NewObjectId()
	newRelease, err := createDefaultTestRelease(targetID)
	if err != nil {
		t.Errorf("[MetaTest] %v", err)
	}

	libraryStore.AddRelease(newRelease)
	// non-existent doc, valid bson ID
	if release, err := libraryStore.GetReleaseByID(bson.NewObjectId()); err == nil {
		t.Errorf("Expected error, did not receive one: '%v'", release)
	}
	// doc exists
	release, err := libraryStore.GetReleaseByID(targetID)
	if err != nil {
		t.Errorf("Error getting doc by id: %v", err)
	}
	if release.KEXPMBID != newRelease.KEXPMBID {
		t.Errorf("MBID fields do not match. Expected: '%v', Actual: '%v'", newRelease.KEXPMBID, release.KEXPMBID)
	}
}