package handlers

import "errors"

const XUser = "X-User"

var ErrInvalidShelfID = errors.New("Invalid Shelf ID")

const headerContentType = "Content-Type"

const contentTypeJSON = "application/json"

const ErrEncodingJSON = "Error converting response value to JSON: "

var ErrInvalidXUser = errors.New("Invalid ID in X-User header. You may have been signed out.")

const ShelvesMineHandlerMethodNotAllowed = "Only allowed to 'GET' from this resource."

const ShelvesHandlerMethodNotAllowed = "Only allowed to 'GET' or 'POST' to this resource."

const ShelfHandlerMethodNotAllowed = "Only allowd to 'GET', 'PATCH', or 'DELETE' to this resource."

const ErrMustBeOwnerToEdit = "You must own the shelf to edit it."

const ErrMustBeOwnerToDelete = "You must own the shelf to delete it."

const DeletedShelfConf = "Deleted shelf\n"
