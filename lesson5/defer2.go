package main

import (
	"fmt"
	"log"
)

type BankAccount struct{}

func (b *BankAccount) Open(accountID string) (*BankAccount, error) {
	fmt.Printf("Opening account with ID: %s\n", accountID)
	return &BankAccount{}, nil
}

func (b *BankAccount) Close() {
	fmt.Println("Closing account")
}

func processAccount(Name_4 string, bank *BankAccount) error {
	fmt.Printf("Processing account with ID: %s\n", Name_4)

	account, err := bank.Open(Name_4)
	if err != nil {
		return err
	}

	defer account.Close()
	fmt.Println("Performing operations on the account")

	return nil
}

func main() {
	// balance := 1000

	// defer fmt.Println("Account balance after transaction: ", balance)

	// balance += 500
	// fmt.Println("Account accessed")

	fmt.Println("Starting the main function")
	bank := &BankAccount{}
	err := processAccount("Account_1234", bank)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Main function completed")
}
