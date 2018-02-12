package shelves

// TODO: Evaluate correct type of 'id' parameters
type ShelvesStore interface {

	// TODO: return array of shelves
	// "GET /v1/shelves/"
	getShelves() error

	// TODO: return single shelf
	// "GET /v1/shelves/{id}"
	getShelfById(id int) error

	// TODO: Will accept a typeof 'Shelf' and replace exisiting shelf
	// at id with that new Shelf
	// "PUT /v1/shelves/{id}"
	updateShelf(id int) error

	// "DELETE /v1/shelves/{id}"
	deleteShelf(id int) error

	// TODO: create and return copy of shelf with given id
	// TODO: Create resource path for this function
	copyShelf(id int) error

	// TODO: export provided shelf as a folder to Dalet
	exportShelf(id int) error
}
