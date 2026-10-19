package main

import (
	"fmt"
	"time"
)

func main() {
	// go processOrder("Buy 1000 shares of APPL")
	// time.Sleep(1 * time.Second)

	go func(order string) {
		fmt.Println("Processing order inside main: ", order)
	}("Buy 1000 shares of MSFT")

	time.Sleep(1 * time.Second)

	fmt.Println("Main function continues executing...")

	orders := []string{"Buy 100 shares of GOOGL", "Sell 50 shares of MSFT", "Buy 200 shares of APPL"}
	ch := make(chan string, len(orders))

	for _, order := range orders {
		go func(o string) {
			ch <- processOrder(o)
		}(order)
	}

	for i := 0; i < len(orders); i++ {
		fmt.Println(<-ch)
	}
}

func processOrder(order string) {
	fmt.Println("Processing order: ", order)
}
