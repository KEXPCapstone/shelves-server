package shelves

// TODO: Evaluate correct type of 'id' parameters
type ShelvesStore interface {
	// TODO: return array of shelves
	getShelves() error

	// TODO: return single shelf
	getShelfById(id int) error

	updateShelf(id int) error

	deleteShelf(id int) error

	// TODO: return copy of shelf with given id
	copyShelf(id int) error

	exportShelf(id int) error
}
