package main

import "fmt"

func main() {
	balance := 1000

	defer fmt.Println("Account balance after transaction: ", balance)

	balance += 500
	fmt.Println("Account accessed")
}
