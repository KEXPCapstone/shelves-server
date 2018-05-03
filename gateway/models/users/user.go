package users

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = 13

//User represents a user account in the database
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email"`
	PassHash  []byte        `json:"-"` //stored, but not encoded to clients
	UserName  string        `json:"userName"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//Updates represents allowed updates to a user profile
type Updates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (nu *NewUser) Validate() error {
	if _, err := mail.ParseAddress(nu.Email); err != nil {
		return fmt.Errorf("%v %v", invalidEmailError, nu.Email)
	}
	if len(nu.UserName) == 0 {
		return errors.New(emptyUserNameError)
	}
	if len(nu.FirstName) == 0 {
		return errors.New(emptyFirstNameError)
	}
	if len(nu.LastName) == 0 {
		return errors.New(emptyLastNameError)
	}
	if len(nu.Password) < 6 {
		return fmt.Errorf("%v %v", newPasswordLengthError, len(nu.Password))
	}
	if nu.Password != nu.PasswordConf {
		return errors.New(newPasswordConfError)
	}
	return nil
}

func (nu *NewUser) ToUser() (*User, error) {
	if err := nu.Validate(); err != nil {
		return nil, err
	}
	email := strings.TrimSpace(nu.Email)
	email = strings.ToLower(email)
	usr := User{
		ID:        bson.NewObjectId(),
		Email:     email,
		UserName:  nu.UserName,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
	}
	if err := usr.SetPassword(nu.Password); err != nil {
		return nil, err
	}
	return &usr, nil
}

func (u *User) FullName() string {
	return strings.Trim(u.FirstName+" "+u.LastName, " ")
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return fmt.Errorf("%v %v", passwordHashError, err)
	}
	u.PassHash = hash
	return nil
}

func (u *User) Authenticate(password string) error {
	if err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password)); err != nil {
		return errors.New(authenticationFailure)
	}
	return nil
}

// TODO: Revisit, see what updates may be relevant.
func (u *User) ApplyUpdates(updates *Updates) error {
	if len(updates.FirstName) == 0 || len(updates.LastName) == 0 {
		return errors.New(invalidNameUpdate)
	}
	u.FirstName = updates.FirstName
	u.LastName = updates.LastName
	return nil
}
