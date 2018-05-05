package releases

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/KEXPCapstone/shelves-server/library/indexes"
	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MongoStore implements LibraryStore interface
type MongoStore struct {
	session           *mgo.Session
	dbname            string
	releaseCollection string
	artistCollection  string
	genreCollection   string
	noteCollection    string
}

// NewMongoStore returns a new MongoStore struct with fields initialized to the passed parameters
func NewMongoStore(sess *mgo.Session, dbName string, releaseCollection string, artistCollection string, genreCollection string, noteCollection string) *MongoStore {
	if sess == nil {
		panic(NoMgoSess)
	}
	return &MongoStore{
		session:           sess,
		dbname:            dbName,
		releaseCollection: releaseCollection,
		artistCollection:  artistCollection,
		genreCollection:   genreCollection,
		noteCollection:    noteCollection,
	}
}

// AddRelease inserts a single release into the library
func (ms *MongoStore) AddRelease(release *Release) (*Release, error) {
	// TODO: Change parameter to something like "New Release", and then call validation methods
	coll := ms.session.DB(ms.dbname).C(ms.releaseCollection)
	if err := coll.Insert(release); err != nil {
		return nil, err
	}
	return release, nil
}

// GetReleases returns releases in the library greater than 'lastID'
// 'limit' specifies the max # of releases to return
func (ms *MongoStore) GetReleases(lastID string, limit int) ([]*Release, error) {
	coll := ms.session.DB(ms.dbname).C(ms.releaseCollection)
	releases := []*Release{}
	if err := coll.Find(bson.M{"_id": bson.M{"$gt": lastID}}).Limit(limit).All(&releases); err != nil {
		return nil, err
	}
	return releases, nil
}

// GetReleaseByID returns a single release in the library
func (ms *MongoStore) GetReleaseByID(id string) (*Release, error) {
	log.Printf("UUID: '%v'", id)
	coll := ms.session.DB(ms.dbname).C(ms.releaseCollection)
	release := &Release{}
	if err := coll.FindId(id).One(release); err != nil {
		log.Print(err)
		return nil, err
	}
	return release, nil
}

// GetReleasesByField accepts a pairing of a field key and value
// returning a slice of releases where release.field['value'] matches the passed params
func (ms *MongoStore) GetReleasesByField(field string, value string) ([]*Release, error) {
	coll := ms.session.DB(ms.dbname).C(ms.releaseCollection)
	releases := []*Release{}
	if err := coll.Find(bson.M{field: value}).All(&releases); err != nil {
		return nil, err
	}
	return releases, nil
}

func (ms *MongoStore) GetReleasesBySliceSearchResults(searchResults []indexes.SearchResult) ([]*ReleaseAndMatchCriteria, error) {
	results := []*ReleaseAndMatchCriteria{}
	for _, match := range searchResults {
		release, err := ms.GetReleaseByID(match.ReleaseID)
		if err != nil {
			return nil, err // should we return this? this would be returned in the case that the object id is in the trie but not in the db...
		}
		results = append(results, &ReleaseAndMatchCriteria{Release: release, IndexInfo: match})
	}

	return results, nil
}

func (ms *MongoStore) IndexReleases() (*indexes.TrieNode, error) {
	coll := ms.session.DB(ms.dbname).C(ms.releaseCollection)
	iter := coll.Find(nil).Iter()
	release := Release{}
	t := indexes.CreateTrieRoot()
	for iter.Next(&release) {
		t.AddToTrie(strings.ToLower(release.KEXPReleaseArtistCredit), indexes.SearchResult{ReleaseID: release.ID, FieldMatchedOn: "KEXPReleaseArtistCredit"})
		t.AddToTrie(strings.ToLower(release.Date), indexes.SearchResult{ReleaseID: release.ID, FieldMatchedOn: "Date"})
		t.AddToTrie(strings.ToLower(release.Title), indexes.SearchResult{ReleaseID: release.ID, FieldMatchedOn: "Title"})
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}
	return t, nil
}

// GetArtists returns artists whose name is alphabetically greater than
// 'lastID'
// 'limit' specifies the max # of docs to return
func (ms *MongoStore) GetArtists(lastID string, limit int) ([]*Artist, error) {
	coll := ms.session.DB(ms.dbname).C(ms.artistCollection)
	artists := []*Artist{}
	collation := &mgo.Collation{Locale: "en", Strength: 1}
	if err := coll.Find(bson.M{"artistName": bson.M{"$gt": lastID}}).Sort("artistName").Collation(collation).Limit(limit).All(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

// GetArtistByMBID returns a specific artist matching the supplied id (MusicBrainz artist MBID)
func (ms *MongoStore) GetArtistByMBID(id string) (*Artist, error) {
	coll := ms.session.DB(ms.dbname).C(ms.artistCollection)
	artist := &Artist{}
	if err := coll.FindId(id).One(artist); err != nil {
		return nil, err
	}
	return artist, nil
}

// GetGenres returns genres in the library greater than 'lastID'
// 'limit' specifies the max # of docs to return
func (ms *MongoStore) GetGenres(lastID bson.ObjectId, limit int) ([]*Genre, error) {
	coll := ms.session.DB(ms.dbname).C(ms.genreCollection)
	genres := []*Genre{}
	if err := coll.Find(bson.M{"_id": bson.M{"$gt": lastID}}).Limit(limit).All(&genres); err != nil {
		return nil, err
	}
	return genres, nil
}

// GetGenreByID returns a specific genre with the supplied id
func (ms *MongoStore) GetGenreByID(id bson.ObjectId) (*Genre, error) {
	coll := ms.session.DB(ms.dbname).C(ms.genreCollection)
	genre := &Genre{}
	if err := coll.FindId(id).One(genre); err != nil {
		return nil, err
	}
	return genre, nil
}

func (ms *MongoStore) AddNoteToRelease(note *Note) (*Note, error) {
	noteColl := ms.session.DB(ms.dbname).C(ms.noteCollection)
	if err := noteColl.Insert(note); err != nil {
		return nil, fmt.Errorf("%v %v", ErrInsertNote, err)
	}
	release, err := ms.GetReleaseByID(note.ReleaseID)
	if err != nil {
		return nil, errors.New(ErrReleaseNotFound)
	}
	release.Notes = append(release.Notes, note.ID)
	releaseColl := ms.session.DB(ms.dbname).C(ms.releaseCollection)
	if err := releaseColl.UpdateId(note.ReleaseID, bson.M{"$set": release}); err != nil {
		return nil, fmt.Errorf("%v %v", ErrAddNoteToRelease, err)
	}
	return note, nil
}

func (ms *MongoStore) GetNotesFromRelease(id string) ([]*Note, error) {
	noteColl := ms.session.DB(ms.dbname).C(ms.noteCollection)
	resultNotes := []*Note{}
	if err := noteColl.Find(bson.M{"releaseID": id}).All(&resultNotes); err != nil {
		return nil, fmt.Errorf("%v %v", ErrFetchingNotes, err)
	}
	return resultNotes, nil
}

// retrieves a list of all distinct artists in the library, sorted alphabetically
// // TODO: refine this query
// // > db.releases.aggregate([{$group: {_id: "$KEXPReleaseArtistCredit", releases: {$push: {id: "$_id", KEXPTitle: "$KEXPTitle", KEXPMBID: "$KEXPMBID"}}}},{$sort:{_id:1}},{$out: "artists"}],{collation:{locale: "en", strength: 1}})
// func (ms *MongoStore) GetAllArtists() ([]string, error) {
// 	db := ms.session.DB(ms.dbname)
// 	result := []string{}
// 	// bson.M for nested fields to be handled gracefully
// 	pipeline := []bson.M{
// 		bson.M{
// 			"$group": bson.M{
// 				"_id":           "$KEXPReleaseArtistCredit",
// 				"totalReleases": bson.M{"$sum": 1},
// 			},
// 		},
// 		bson.M{
// 			"$sort": bson.M{"_id": 1},
// 		},
// 	}
// 	collation := bson.M{"locale": "en", "strength": 1}
// 	// using bson.D here to maintain ordering, since mongo will interpret first key as the command name
// 	db.Run(bson.D{{"aggregate", "releases"}, {"pipeline", pipeline}, {"collation", collation}}, &result)
// 	if err := db.Run(bson.D{{"aggregate", "releases"}, {"pipeline", pipeline}, {"collation", collation}}, &result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
