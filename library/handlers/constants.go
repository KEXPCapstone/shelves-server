package handlers

const ErrInsertRelease = "Error inserting release: "

const ErrFetchingRelease = "Error fetching release(s): "

const ErrFetchingArtist = "Error fetching artist(s): "

const ErrInvalidReleaseID = "Invalid Release ID"

const ErrDecodingJSON = "Error decoding JSON: "

const ErrEncodingJSON = "Error converting response value to JSON: "

const HandlerInvalidMethod = "%v requests are not allowed for this resource"

const headerContentType = "Content-Type"

const contentTypeJSON = "application/json"

const maxSearchResults = 100

const ErrorSearching = "Error processing release results: "

const XUser = "X-User"

const ErrInvalidXUser = "Invalid ID in X-User header. You may have been signed out."

const ErrNoXUser = "Invalid or no X-User header. You may have been signed out."
