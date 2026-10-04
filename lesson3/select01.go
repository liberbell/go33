package main

import (
	"fmt"
	"time"
)

func main() {
	bankingOperations := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		bankingOperations <- "Deposit complete"
	}()

	select {
	case msg := <-bankingOperations:
		fmt.Println(msg)
	}
}
