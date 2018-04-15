package releases

import (
	"strings"

	"github.com/KEXPCapstone/shelves-server/library/indexes"
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

func (ms *MgoStore) Insert(release *Release) error {
	// TODO: Change parameter to something like "New Release", and then call validation methods
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	if err := coll.Insert(release); err != nil {
		return err
	}
	return nil
}

func (ms *MgoStore) GetAllReleases() ([]*Release, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	releases := []*Release{}
	if err := coll.Find(nil).All(&releases); err != nil {
		return nil, err
	}
	return releases, nil
}

func (ms *MgoStore) GetReleaseByID(id bson.ObjectId) (*Release, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	release := &Release{}
	if err := coll.FindId(id).One(release); err != nil {
		return nil, err
	}
	return release, nil
}

func (ms *MgoStore) GetReleasesByField(field string, value string) ([]*Release, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	releases := []*Release{}
	if err := coll.Find(bson.M{field: value}).All(&releases); err != nil {
		return nil, err
	}
	return releases, nil
}

func (ms *MgoStore) GetReleasesBySliceSearchResults(searchResults []indexes.SearchResult) ([]*ReleaseAndMatchCriteria, error) {
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

func (ms *MgoStore) IndexReleases() (*indexes.TrieNode, error) {
	coll := ms.session.DB(ms.dbname).C(ms.colname)
	iter := coll.Find(nil).Iter()
	release := Release{}
	t := indexes.CreateTrieRoot()
	for iter.Next(&release) {
		t.AddToTrie(strings.ToLower(release.KEXPReleaseArtistCredit), indexes.SearchResult{ReleaseID: release.ID, FieldMatchedOn: "KEXPReleaseArtistCredit"})
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}
	return t, nil
}
