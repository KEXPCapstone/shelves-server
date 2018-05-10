package handlers

import "errors"

const XUser = "X-User"

var ErrInvalidShelfID = errors.New("Invalid Shelf ID")

const headerContentType = "Content-Type"

const contentTypeJSON = "application/json"

const ErrEncodingJSON = "Error converting response value to JSON: "

const ErrDecodingJSON = "Error decoding request body: "

var ErrInvalidXUser = errors.New("Invalid ID in X-User header. You may have been signed out.")

const ShelvesMineHandlerMethodNotAllowed = "Only allowed to 'GET' from this resource."

const ShelvesHandlerMethodNotAllowed = "Only allowed to 'GET' or 'POST' to this resource."

const ShelfHandlerMethodNotAllowed = "Only allowd to 'GET', 'PUT', or 'DELETE' to this resource."

const ErrMustBeOwnerToEdit = "You must own the shelf to edit it."

const ErrMustBeOwnerToDelete = "You must own the shelf to delete it."

const DeletedShelfConf = "Successfully deleted shelf\n"

const ErrUpdateShelf = "Error encountered updating shelf"

const UpdatedShelfConf = "Successfully Updated shelf\n"

const UserShelvesHandlerMethodNotAllowed = "Only allowed to 'GET' from this resource."

const InvalidUserID = "Invalid user ID in request URL."

const FeaturedShelvesHandlerMethodNotAllowed = "Only allowed to 'GET' from this resource."

const ErrNoXUser = "Invalid or no X-User header. You may have been signed out."