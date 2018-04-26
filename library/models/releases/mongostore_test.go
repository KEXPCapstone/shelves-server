package releases

import (
	"encoding/json"
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"
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
	noteCol := "notes"
	ms := NewMongoStore(mongoSess, dbname, releaseCol, artistCol, genreCol, noteCol)
	return ms, nil
}

func TestNewMongoStore(t *testing.T) {
	dbname := "library"
	releaseCol := "releases"
	artistCol := "artists"
	genreCol := "genres"
	noteCol := "notes"

	mongoSess, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error dialing mongodb: %v", err)
	}
	libraryStore := NewMongoStore(mongoSess, dbname, releaseCol, artistCol, genreCol, noteCol)
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
	if libraryStore.noteCollection != noteCol {
		t.Errorf("Error setting note collection name. Expected: '%v' Actual: '%v'", noteCol, libraryStore.noteCollection)
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
	_ = NewMongoStore(nil, "foo", "bar", "baz", "bip", "woop")
}

func TestAddRelease(t *testing.T) {
	libraryStore, err := createTestingMgoStore()
	if err != nil {
		t.Errorf("[MetaTest] Error creating test MongoStore: '%v", err)
	}
	newRelease, err := createDefaultTestRelease()
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
func createDefaultTestRelease() (*Release, error) {
	// unmarshal an example doc into release struct
	mockReleaseJSON := `{"id":"53042259-1287-4f47-9a99-5a7413df7b3f","artistCredit":[{"artist":{"disambiguation":"∆","id":"fc7bbf00-fbaa-4736-986b-b3ac0266ca9b","name":"alt-J","sort-name":"alt-J"},"joinphrase":"","name":"alt-J"}],"releaseEvents":[{"area":{"disambiguation":"","id":"106e0bec-b638-3b37-b731-f53d507dc00e","iso-3166-1-codes":["AU"],"name":"Australia","sort-name":"Australia"},"date":"2012"}],"coverArtArchive":{"artwork":true,"back":false,"count":1,"darkened":false,"front":true},"KEXPReleaseGroupMBID":"0d8562eb-7f72-427b-8a0b-984cc5ee7766","KEXPReleaseArtistCredit":"alt-J","labelInfo":[{"catalog-number":"LIB140CD","label":{"disambiguation":"","id":"8c63a604-872c-488e-af18-5cead3d82f17","label-code":null,"name":"Liberator Music","sort-name":"Liberator Music"}}],"media":[],"Status":"Official","Disambiguation":"","barcode":"9341004016156","packaging":"Jewel Case","date":"2012","asin":"","KEXPPrimaryGenre":"Rock/Pop","title":"An Awesome Wave","countryCode":"AU","yellows":null,"reds":null}`
	newRelease := &Release{}
	buffer := []byte(mockReleaseJSON)
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
	newRelease, err := createDefaultTestRelease()
	if err != nil {
		t.Errorf("[MetaTest] %v", err)
	}

	libraryStore.AddRelease(newRelease)
	// non-existent doc, valid bson ID
	if release, err := libraryStore.GetReleaseByID("not a real ID"); err == nil {
		t.Errorf("Expected error, did not receive one: '%v'", release)
	}
	// doc exists
	release, err := libraryStore.GetReleaseByID("53042259-1287-4f47-9a99-5a7413df7b3f")
	if err != nil {
		t.Errorf("Error getting doc from id '%v': '%v'", newRelease.ID, err)
	}
	if release.ID != newRelease.ID {
		t.Errorf("MBID fields do not match. Expected: '%v', Actual: '%v'", newRelease.ID, release.ID)
	}
}
