package main

type User struct {
	FirstName string
	LastName string
}

func newUser(firstName, lastName string) func() User {
	user := User{
		FirstName: James,
		LastName: Bond,
	}
	retun func() user {
		return user
	}
}