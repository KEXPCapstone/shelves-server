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
	AddRelease(release *Release) (*Release, error)

	// return all releases in the library
	GetReleases(lastID string, limit int) ([]*Release, error)

	// return a single release with the supplied id
	GetReleaseByID(id string) (*Release, error)

	// return a slice of releases matching the given field value
	GetReleasesByField(field string, value string) ([]*Release, error)

	// return all artists in the library
	GetArtists(lastID string, limit int) ([]*Artist, error)

	// return a specific artist with the supplied musicbrainz artist MBID
	GetArtistByMBID(id string) (*Artist, error)

	// return all genres in the library
	GetGenres(lastID bson.ObjectId, limit int) ([]*Genre, error)

	// return a specific genre with the supplied id
	GetGenreByID(id bson.ObjectId) (*Genre, error)

	GetReleasesBySliceSearchResults(searchResults []indexes.SearchResult) ([]*ReleaseAndMatchCriteria, error)

	IndexReleases() (*indexes.TrieNode, error)

	AddNoteToRelease(releaseID string, note *Note) (*Note, error)
}
