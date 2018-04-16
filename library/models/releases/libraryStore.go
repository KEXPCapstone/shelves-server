package releases

// TODO: general interface should be agnostic of mgo types
import (
	"github.com/KEXPCapstone/shelves-server/library/indexes"
	"gopkg.in/mgo.v2/bson"
)

// LibraryStore provides functions for interacting with a database
// representing the KEXP music library
type LibraryStore interface {

	// add a single release into the library
	AddRelease(release *Release) error

	// return all releases in the library
	GetReleases() ([]*Release, error)

	// return a single release with the supplied id
	GetReleaseByID(id bson.ObjectId) (*Release, error)

	// return a slice of releases matching the given field value
	GetReleasesByField(field string, value string) ([]*Release, error)

	// return all artists in the library
	GetArtists() ([]*Artist, error)

	// return a specific artist with the supplied musicbrainz artist MBID
	GetArtistByMBID(id string) (*Artist, error)

	// return all genres in the library
	GetGenres() ([]*Genre, error)

	// return a specific genre with the supplied id
	GetGenreByID(id bson.ObjectId) (*Genre, error)

	GetReleasesBySliceSearchResults(searchResults []indexes.SearchResult) ([]*ReleaseAndMatchCriteria, error)

	IndexReleases() (*indexes.TrieNode, error)
}
