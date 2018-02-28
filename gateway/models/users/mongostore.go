package users

// implements UserStore interface
type MgoStore struct {
}

func (ms *MgoStore) Insert(nu *NewUser) (*User, error) {
	// TODO: Convert nu to an "intermediate user"
	// Place user into db, returning the associated id
	// Add the id field into the user?

	// Alternatively, convert nu to User without a
	// value for id, only pass in the relevant fields to be
	// added to a row in the database, and then write the
	// returned id to the User
	return nil, nil
}

func (ms *MgoStore) GetByID(id int) (*User, error) {
	return nil, nil
}

func (ms *MgoStore) GetByEmail(email string) (*User, error) {
	return nil, nil
}

func (ms *MgoStore) GetByUserName(username string) (*User, error) {
	return nil, nil
}

func (ms *MgoStore) Update(userID int, updates *Updates) error {
	return nil
}

func (ms *MgoStore) Delete(userID int) error {
	return nil
}
