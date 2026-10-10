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

func main() {
	deposit := newAccount(100.50)
	fmt.Println(deposit(80.0))
	fmt.Println(deposit(1100))
}
