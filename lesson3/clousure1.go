package main

type User struct {
	FirstName string
	LastName  string
}

func newUser(firstName, lastName string) func() User {
	user := User{
		FirstName: firstName,
		LastName:  lastName,
	}
	return func() User {
		return user
	}
}
