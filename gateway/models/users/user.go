package users

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

func (nu *NewUser) Validate() error {}

func (nu *NewUser) ToUser() (*User, error) {}

func (u *User) FullName() string {}

func (u *User) SetPassword(password string) error {}

func (u *User) Authenticate(password string) error {}

func (u *User) ApplyUpdates(updates *Updates) error {}
