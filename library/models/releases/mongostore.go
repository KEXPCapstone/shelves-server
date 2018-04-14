package releases

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

// retrieves a list of all distinct artists in the library, sorted alphabetically
// TODO: implement this query:
// db.releases.aggregate([{$group: {_id: "$KEXPReleaseArtistCredit", totalReleases: {$sum:1}}},{$sort:{_id: 1}}], {collation:{locale: "en", strength: 1}})
func (ms *MgoStore) GetAllArtists() ([]string, error) {
	db := ms.session.DB(ms.dbname)
	result := []string{}
	// bson.M for nested fields to be handled gracefully
	pipeline := []bson.M{
		bson.M{
			"$group": bson.M{
				"_id":           "$KEXPReleaseArtistCredit",
				"totalReleases": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{"_id": 1},
		},
	}
	collation := bson.M{"locale": "en", "strength": 1}
	// using bson.D here to maintain ordering, since mongo will interpret first key as the command name
	db.Run(bson.D{{"aggregate", "releases"}, {"pipeline", pipeline}, {"collation", collation}}, &result)
	if err := db.Run(bson.D{{"aggregate", "releases"}, {"pipeline", pipeline}, {"collation", collation}}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
