package handlers

import "errors"

const XUser = "X-User"

const ErrInvalidShelfID = errors.New("Invalid Shelf ID")

const headerContentType = "Content-Type"

const contentTypeJSON = "application/json"

const ErrEncodingJSON = "Error converting response value to JSON: "
