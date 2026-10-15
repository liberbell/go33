package main

import (
	"database/sql"
	"fmt"
)

func getUserProfile(Name_1 string, db *sql.DB) {
	fmt.Printf("Attempting to get user profile for: %s\n", Name_1)

	fmt.Println("Retriving user profile...")
}

func main() {
	defer fmt.Println("Withdrawal completed")
	defer fmt.Println("Deposit completed")
	fmt.Println("Account balance checked")

}
