package main

import "fmt"

func processTransaction(amount float64, operation func(float64) float64) float64 {
	return operation(amount)
}

func main() {
	func(firstName, lastName, email string) {
		fmt.Printf("User: %s %s, Email: %s\n", firstName, lastName, email)
	}("James", "Bond", "james@bond.com")

	deposit := func(amount float64) float64 {
		fmt.Printf("Deposited: $%.2f\n", amount)
		return amount
	}

	newBalance := processTransaction(500.50, deposit)
	fmt.Printf("New Balance: $%.2f\n", newBalance)
}
