package main

import "fmt"

func main() {
	defer fmt.Println("Withdrawal completed")
	defer fmt.Println("Deposit completed")
}
