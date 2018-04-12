package handlers

const ErrInsertRelease = "Error inserting release: "

const ErrFetchingRelease = "Error fetching release(s): "

const ErrInvalidReleaseID = "Invalid Release ID"

const ErrDecodingJSON = "Error decoding JSON: "

const ErrEncodingJSON = "Error converting response value to JSON: "

const ReleasesHandlerInvalidMethod = "Only POST and GET requests are allowed for this resource"

const SingleReleaseHandlerInvalidMethod = "Only GET requests are allowed for this resource"

const headerContentType = "Content-Type"

const contentTypeJSON = "application/json"

const maxSearchResults = 100

const ErrorSearching = "Error processing release results: "
