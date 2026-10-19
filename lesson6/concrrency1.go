package main

import (
	"fmt"
	"time"
)

func main() {
	go processOrder("Buy 1000 shares of APPL")
	time.Sleep(1 * time.Second)
}

func processOrder(order string) {
	fmt.Println("Processing order: ", order)
}
