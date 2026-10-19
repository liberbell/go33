package main

import (
	"fmt"
	"sync"
)

func main() {

	var waitgroup sync.WaitGroup

	orders := []string{"Buying 100 shares of APPL", "Selling 50 shares of GOOGL", "Buying 200 shares of MSFT"}
	waitgroup.Add(len(orders))

	for _, order := range orders {
		go func(o string) {
			defer waitgroup.Done()
			fmt.Println("Processing order: ", o)
		}(order)
	}

	waitgroup.Wait()
	fmt.Println("All orders have been processed")
}
