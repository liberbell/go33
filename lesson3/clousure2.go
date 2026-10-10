package main

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
