package main

import "fmt"

func main() {
	func(firstName, lastName, email string) {
		fmt.Printf("User: %s %s, Email: %s", firstName, lastName, email)
	}("James", "Bond", "james@bond.com")
}
