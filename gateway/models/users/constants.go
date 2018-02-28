package users

import "errors"

const invalidEmailError = "Email must be a valid email address. Email provided: "

const emptyUserNameError = "UserName cannot be empty."

const emptyFirstNameError = "First name must be at least 1 character"

const emptyLastNameError = "Last name must be at least 1 character"

const newPasswordLengthError = "Password must be at least 6 characters.  Length provided: "

const newPasswordConfError = "Password confirmation does not match password"

const passwordHashError = "Error hashing password: "

const authenticationFailure = "Failed to authenticate; incorrect password."

const invalidNameUpdate = "Invalid update: first name and last name must be at least 1 character"

var ErrUserNotFound = errors.New("user not found")

var ErrInsertUser = errors.New("Error inserting user into DB")

var NoMgoSess = "nil pointer passed from session"

var ErrStrUpdateUser = "Error returned when updating user"
