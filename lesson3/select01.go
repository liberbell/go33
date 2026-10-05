package main

import (
	"fmt"
	"time"
)

func main() {
	bankingOperations := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		bankingOperations <- "Deposit complete"
	}()

	select {
	case msg := <-bankingOperations:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Operation timeout")
	}
}
