package users

import (
	"errors"
	"fmt"
	"net/mail"
)

var bcryptCost = 13

//User represents a user account in the database
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	PassHash  []byte `json:"-"` //stored, but not encoded to clients
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
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
		return fmt.Errorf("Email must be a valid email address. Email provided: %v", nu.Email)
	}
	if len(nu.UserName) == 0 {
		return fmt.Errorf("UserName cannot be empty. Username provided: %v", nu.UserName)
	}
	if len(nu.FirstName) == 0 {
		return errors.New("First name must be at least 1 character")
	}
	if len(nu.LastName) == 0 {
		return errors.New("Last name must be at least 1 character")
	}
	if len(nu.Password) < 6 {
		return fmt.Errorf("Password must be at least 6 characters.  Length provided: %v", len(nu.Password))
	}
	if nu.Password != nu.PasswordConf {
		return errors.New("Password confirmation does not match password")
	}
	return nil
}

func (nu *NewUser) ToUser() (*User, error) {}

func (u *User) FullName() string {}

func (u *User) SetPassword(password string) error {}

func (u *User) Authenticate(password string) error {}

func (u *User) ApplyUpdates(updates *Updates) error {}
