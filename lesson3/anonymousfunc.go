package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello from an annonymous function")
	}()

	func(firstName, lastName string) {
		fmt.Printf("User profile: %s %s\n", firstName, lastName)
	}("FirstName", "LastName")

	displayBalance := func(accountHolder, string, balance float64) {

	}
}
