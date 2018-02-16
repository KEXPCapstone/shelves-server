package users

type UserStore interface {
	//GetByID returns the User with the given ID
	GetByID(id int) (*User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*User, error)

	//Insert converts the NewUser to a User, inserts
	//it into the database, and returns it
	Insert(newUser *NewUser) (*User, error)

	//Update applies UserUpdates to the given user ID
	Update(userID int, updates *Updates) error

	//Delete deletes the user with the given ID
	Delete(userID int) error
}
