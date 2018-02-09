package sessions

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"io"
)

// TODO: I received some feeback on the assignment this is taken from
// on certain changes to make. See https://github.com/info344-a17/challenges-abourn/pull/3
// for potential improvements

//InvalidSessionID represents an empty, invalid session ID
const InvalidSessionID SessionID = ""

//idLength is the length of the ID portion
const idLength = 32

//signedLength is the full length of the signed session ID
//(ID portion plus signature)
const signedLength = idLength + sha256.Size

//SessionID represents a valid, digitally-signed session ID.
//This is a base64 URL encoded string created from a byte slice
//where the first `idLength` bytes are crytographically random
//bytes representing the unique session ID, and the remaining bytes
//are an HMAC hash of those ID bytes (i.e., a digital signature).
//The byte slice layout is like so:
//+-----------------------------------------------------+
//|...32 crypto random bytes...|HMAC hash of those bytes|
//+-----------------------------------------------------+
type SessionID string

//ErrInvalidID is returned when an invalid session id is passed to ValidateID()
var ErrInvalidID = errors.New("Invalid Session ID")

//NewSessionID creates and returns a new digitally-signed session ID,
//using `signingKey` as the HMAC signing key. An error is returned only
//if there was an error generating random bytes for the session ID
func NewSessionID(signingKey string) (SessionID, error) {
	if len(signingKey) == 0 {
		return InvalidSessionID, errors.New("Signing key cannot be empty")
	}

	r := make([]byte, idLength)
	_, err := rand.Read(r)
	if err != nil {
		return InvalidSessionID, errors.New("Error generating random ")
	}
	h := hmac.New(sha256.New, []byte(signingKey))
	if _, err := io.Copy(h, bytes.NewReader(r)); err != nil {
		return InvalidSessionID, errors.New("Error hashing byte")
	}
	hash := h.Sum(nil)
	resultSlice := append([]byte{}, r...)
	resultSlice = append(resultSlice, hash...)
	sid := base64.URLEncoding.EncodeToString(resultSlice)
	return SessionID(sid), nil
}

//ValidateID validates the string in the `id` parameter
//using the `signingKey` as the HMAC signing key
//and returns an error if invalid, or a SessionID if valid
func ValidateID(id string, signingKey string) (SessionID, error) {

	sid, err := base64.URLEncoding.DecodeString(id)
	if err != nil {
		return InvalidSessionID, ErrInvalidID
	}
	h := hmac.New(sha256.New, []byte(signingKey))
	if _, err := io.Copy(h, bytes.NewReader(sid[:idLength])); err != nil {
		return InvalidSessionID, errors.New("Error hashing")
	}
	hash := h.Sum(nil)
	if subtle.ConstantTimeCompare(hash, sid[idLength:]) == 1 {
		return SessionID(id), nil
	}
	return InvalidSessionID, ErrInvalidID
}

//String returns a string representation of the sessionID
func (sid SessionID) String() string {
	return string(sid)
}
