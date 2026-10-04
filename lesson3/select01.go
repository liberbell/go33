package main

import "time"

func main() {
	bankingOperations := make(chan stirng)

	go func() {
		time.Sleep(2 * time.Second)
		bankingOperations <- "Deposit complete"
	}()
}
