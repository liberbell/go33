package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Bank struct{}

func getUserProfile(Name_1 string, db *sql.DB) {
	fmt.Printf("Attempting to get user profile for: %s\n", Name_1)

	fmt.Println("Retriving user profile...")
}

func (b *Bank) Open(accountName string) (*Bank, error) {
	fmt.Printf("Opening account: %s\n", accountName)
	return &Bank{}, nil
}

func (b *Bank) Close() {
	fmt.Println("Closing account")
}

func (b *Bank) Transfer(source, destination *Bank) error {
	fmt.Println("Transfering funds from source to destination account")
	return nil
}

func transferFunds(Name_2, Name_3 string, bank *Bank) error {
	source, err := bank.Open(Name_2)
	if err != nil {
		return err
	}

	defer func() {
		source.Close()
		fmt.Printf("Account %s closed\n", Name_2)
	}()

	destination, err := bank.Open(Name_3)
	if err != nil {
		return err
	}

	defer func() {
		destination.Close()
		fmt.Printf("Account %s closed\n", Name_3)
	}()

	err = bank.Transfer(source, destination)
	if err != nil {
		return err
	}
}

func main() {
	defer fmt.Println("Withdrawal completed")
	defer fmt.Println("Deposit completed")
	fmt.Println("Account balance checked")

	fmt.Println("Starting the main function")

	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		db.Close()
		fmt.Println("Database connection closed")
	}()

	getUserProfile("Apple", db)
	fmt.Println("Main function completed")

}
