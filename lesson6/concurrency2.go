package main

import (
	"fmt"
	"sync"
)

type OrderBook struct {
	mx     sync.Mutex
	orders []string
}

func (ob *OrderBook) addOrder(wg *sync.WaitGroup, order string) {
	ob.mx.Lock()
	defer ob.mx.Unlock()
	defer wg.Done()

	ob.orders = append(ob.orders, order)
	fmt.Println("Added order: ", order)
}

func main() {

	// var waitgroup sync.WaitGroup

	// orders := []string{"Buying 100 shares of APPL", "Selling 50 shares of GOOGL", "Buying 200 shares of MSFT"}
	// waitgroup.Add(len(orders))

	// for _, order := range orders {
	// 	go func(o string) {
	// 		defer waitgroup.Done()
	// 		fmt.Println("Processing order: ", o)
	// 	}(order)
	// }

	// waitgroup.Wait()
	// fmt.Println("All orders have been processed")

	var wg sync.WaitGroup
	ob := OrderBook{}

	orders := {"Buy 100 shares of APPL", "Sell 50 shares of MSFT", "Buy 200 shares of GOOGLE"}

	wg.Add(len(orders))

	for _, order := range orders {
		go ob.addOrder(&wg, order)
	}

	wg.Wait()
	fmt.Println("All orders have been added:", ob.orders)

}
