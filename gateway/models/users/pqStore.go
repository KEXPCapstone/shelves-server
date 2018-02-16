package users

// implements UserStore interface
type PqStore struct {
}

func (ps *PqStore) Insert(nu *NewUser) (*User, error) {
	// TODO: Convert nu to an "intermediate user"
	// Place user into db, returning the associated id
	// Add the id field into the user?

	// Alternatively, convert nu to User without a
	// value for id, only pass in the relevant fields to be
	// added to a row in the database, and then write the
	// returned id to the User
}

func (ps *PqStore) GetByID(id int) (*User, error) {}

func (ps *PqStore) GetByEmail(email string) (*User, error) {}

func (ps *PqStore) GetByUserName(username string) (*User, error) {}

func (ps *PqStore) Update(userID int, updates *Updates) error {}

func (ps *PqStore) Delete(userID int) error {}
