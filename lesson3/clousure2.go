package main

import "fmt"

type BankAccount struct {
	Balance float64
}

func newAccount(initialBalance float64) func(float64) float64 {
	account := BankAccount{Balance: initialBalance}
	return func(amount float64) float64 {
		account.Balance += amount
		return account.Balance
	}
}

func accountOperations() func(float64) float64 {
	balance := 0.0
	return func(amount float64) float64 {
		balance += amount
		return balance
	}
}

func main() {
	deposit := newAccount(100.50)
	fmt.Println(deposit(80.0))
	fmt.Println(deposit(1100))
}
