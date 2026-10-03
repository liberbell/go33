package main

import "fmt"

func main() {
	const (
		CheckingAccount = iota
		SavingAccount
		BusinessAccount
	)

	fmt.Println("CheckingAccount:", CheckingAccount)
	fmt.Println("SavingAccount:", SavingAccount)
	fmt.Println("BusinessAccount:", BusinessAccount)
}
