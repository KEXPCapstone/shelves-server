package users

import (
	"fmt"
	"reflect"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// a helper function which will create a new MongoDB database
// and collection.  Returns a pointer to the associated MgoStore or an error
func createTestingMgoStore() (*MgoStore, error) {
	mongoSess, err := mgo.Dial("localhost")
	if err != nil {
		return nil, fmt.Errorf("Error dialing MongoDB. Error: %v", err)
	}
	ms := NewMgoStore(mongoSess, "users", "users")
	return ms, nil
}

// a helper function for testing which inserts a valid user object
// into the provided MongoDB database
func insertTestingData(ms *MgoStore) (*User, error) {
	nu := NewUser{
		Email:        "test@example.com",
		Password:     "password",
		PasswordConf: "password",
		UserName:     "uname",
		FirstName:    "fname",
		LastName:     "lname",
	}
	usr, err := ms.Insert(&nu)
	if err != nil {
		return nil, fmt.Errorf("Error inserting user into DB: %v", err)
	}
	return usr, nil
}

func TestNewMgoStore(t *testing.T) {
	mongoSess, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error dialing mongodb. Error: %v", err)
	}
	mgoStore := NewMgoStore(mongoSess, "users", "users")
	if mgoStore.session != mongoSess {
		t.Errorf("Error setting session. Expected: %v Actual: %v", mongoSess, mgoStore.session)
	}
	if mgoStore.dbname != "users" {
		t.Errorf("Error setting database name.  Expected: 'users' Actual: %v", mgoStore.dbname)
	}
	if mgoStore.colname != "users" {
		t.Errorf("Error settign collection name. Expected: 'users' Actual: %v", mgoStore.colname)
	}
}

func TestMgoGetByID(t *testing.T) {
	mgoStore, err := createTestingMgoStore()
	if err != nil {
		t.Error(err)
	}
	usr, err := insertTestingData(mgoStore)
	if err != nil {
		t.Errorf("Error inserting test user into DB: %v", err)
	}
	cases := []struct {
		name        string
		id          bson.ObjectId
		expectedErr bool
	}{
		{
			"Valid lookup by ID",
			usr.ID,
			false,
		},
		{
			"Invalid lookup by ID",
			bson.NewObjectId(),
			true,
		},
	}
	for _, c := range cases {
		dbUsr, err := mgoStore.GetByID(c.id)
		if err != nil && !c.expectedErr {
			t.Errorf("Case name: %v. Unexpected error on lookup: %v", c.name, err)
		}
		if err == nil && c.expectedErr {
			t.Errorf("Case name: %v. Expected error but none found. ", c.name)
		}
		if !reflect.DeepEqual(dbUsr, usr) && !c.expectedErr {
			t.Errorf("Case name: %v. User returned does not match expected. Returned: %v Expected: %v", c.name, usr, dbUsr)
		}
	}
}

func TestMgoGetByEmail(t *testing.T) {
	mgoStore, err := createTestingMgoStore()
	if err != nil {
		t.Error(err)
	}
	usr, err := insertTestingData(mgoStore)
	if err != nil {
		t.Errorf("Error inserting test user into DB: %v", err)
	}
	cases := []struct {
		name        string
		email       string
		expectedErr bool
	}{
		{
			"Valid Lookup by Email",
			usr.Email,
			false,
		},
		{
			"Invalid Lookup by Email",
			"",
			true,
		},
	}
	for _, c := range cases {
		dbUsr, err := mgoStore.GetByEmail(c.email)
		if err != nil && !c.expectedErr {
			t.Errorf("Case name: %v. Received unexpected error: %v", c.name, err)
		}
		if err == nil && c.expectedErr {
			t.Errorf("Case name: %v. Expected to find error but none found", c.name)
		}
		if !c.expectedErr && (dbUsr.UserName != usr.UserName || dbUsr.Email != usr.Email) {
			t.Errorf("User in database does not match expected.  Expected: %v Found: %v", usr, dbUsr)
		}
	}
}

func TestMgoGetByUserName(t *testing.T) {
	mgoStore, err := createTestingMgoStore()
	if err != nil {
		t.Error(err)
	}
	usr, err := insertTestingData(mgoStore)
	if err != nil {
		t.Errorf("Error inserting test user into DB: %v", err)
	}
	cases := []struct {
		name        string
		username    string
		expectedErr bool
	}{
		{
			"Valid Lookup by Email",
			usr.UserName,
			false,
		},
		{
			"Invalid Lookup by Email",
			"",
			true,
		},
	}
	for _, c := range cases {
		dbUsr, err := mgoStore.GetByUserName(c.username)
		if err != nil && !c.expectedErr {
			t.Errorf("Case name: %v. Received unexpected error: %v", c.name, err)
		}
		if err == nil && c.expectedErr {
			t.Errorf("Case name: %v. Expected to find error but none found", c.name)
		}
		if !c.expectedErr && (dbUsr.UserName != usr.UserName || dbUsr.Email != usr.Email) {
			t.Errorf("User in database does not match expected.  Expected: %v Found: %v", usr, dbUsr)
		}
	}

}

func TestMgoInsert(t *testing.T) {
	mgoStore, err := createTestingMgoStore()
	if err != nil {
		t.Error(err)
	}
	cases := []struct {
		name        string
		nu          NewUser
		expectedErr bool
	}{
		{
			"Valid user Insertion",
			NewUser{
				Email:        "test@example.com",
				Password:     "password",
				PasswordConf: "password",
				UserName:     "uname",
				FirstName:    "fname",
				LastName:     "lname",
			},
			false,
		},
		{
			"Invalid user Insertion",
			NewUser{
				Email:        "email",
				Password:     "pass",
				PasswordConf: "word",
				UserName:     "user",
			},
			true,
		},
	}
	for _, c := range cases {
		usr, err := mgoStore.Insert(&c.nu)
		if err != nil && !c.expectedErr {
			t.Errorf("Error found when trying to insert user into database. Error found: %v", err)
		}
		if usr != nil { // if current case is for a valid insertion
			dbUsr, err := mgoStore.GetByID(usr.ID)
			if usr.Email != c.nu.Email || usr.FirstName != c.nu.FirstName || usr.LastName != c.nu.LastName || usr.UserName != c.nu.UserName {
				t.Errorf("Returned user does not match expected. User Returned: %v New User Given: %v", usr, c.nu)
			}
			if err != nil && !c.expectedErr {
				t.Errorf("User not inserted into database. Lookup returned error: %v", err)
			}
			if !reflect.DeepEqual(usr, dbUsr) && !c.expectedErr {
				t.Errorf("User inserted into database does not match user returned. User returned: %v User in DB: %v", usr, dbUsr)
			}
		}
	}
}

func TestMgoUpdate(t *testing.T) {
	mgoStore, err := createTestingMgoStore()
	if err != nil {
		t.Error(err)
	}
	usr, err := insertTestingData(mgoStore)
	if err != nil {
		t.Errorf("Error inserting test user into DB: %v", err)
	}
	cases := []struct {
		name        string
		id          bson.ObjectId
		upd         Updates
		expectedErr bool
	}{
		{
			"Valid updates",
			usr.ID,
			Updates{FirstName: "fname2", LastName: "lname2"},
			false,
		},
		{
			"Invalid updates",
			usr.ID,
			Updates{},
			true,
		},
		{
			"Attempt to update non-existent user",
			bson.NewObjectId(),
			Updates{FirstName: "fname2", LastName: "lname2"},
			true,
		},
	}
	for _, c := range cases {
		err := mgoStore.Update(c.id, &c.upd)
		if err != nil && !c.expectedErr {
			t.Errorf("Case name: %v. Encountered unexpected error in updating: %v", c.name, err)
		}
		if err == nil && c.expectedErr {
			t.Errorf("Case name: %v. Expected error but none found", c.name)
		}
		if err == nil && !c.expectedErr {
			usr, err := mgoStore.GetByID(c.id)
			if err != nil {
				t.Errorf("Could not find user after update. Received error: %v", err)
			}
			if usr.FirstName != c.upd.FirstName || usr.LastName != c.upd.LastName {
				t.Error("User was not updated in MongoDB.")
			}
		}
	}
}

func TestMgoDelete(t *testing.T) {
	mgoStore, err := createTestingMgoStore()
	if err != nil {
		t.Error(err)
	}
	usr, err := insertTestingData(mgoStore)
	if err != nil {
		t.Errorf("Error inserting test user into DB: %v", err)
	}
	if err := mgoStore.Delete(usr.ID); err != nil {
		t.Errorf("Expected deletion but instead couldn't find user. Error: %v", err)
	}
	if _, err := mgoStore.GetByID(usr.ID); err == nil {
		t.Error("Test user not actually deleted")
	}
	if err := mgoStore.Delete(bson.NewObjectId()); err == nil {
		t.Error("Expected to not find user but instead deleted another user.")
	}
}
