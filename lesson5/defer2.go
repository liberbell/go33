package main

import "fmt"

type BankAccount struct{}

func (b *BankAccount) Open(accountID string) (*BankAccount, error) {
	fmt.Printf("Opening account with ID: %s\n", accountID)
	return &BankAccount{}, nil
}

func main() {
	balance := 1000

	defer fmt.Println("Account balance after transaction: ", balance)

	balance += 500
	fmt.Println("Account accessed")
}
