package main

import "fmt"

func processTransaction(amount float64, operation func(float64) float64) float64 {
	return operation(amount)
}

func main() {
	func(firstName, lastName, email string) {
		fmt.Printf("User: %s %s, Email: %s", firstName, lastName, email)
	}("James", "Bond", "james@bond.com")
}
