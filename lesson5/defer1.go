package main

import (
	"database/sql"
	"fmt"
	"log"
)

func getUserProfile(Name_1 string, db *sql.DB) {
	fmt.Printf("Attempting to get user profile for: %s\n", Name_1)

	fmt.Println("Retriving user profile...")
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

}
