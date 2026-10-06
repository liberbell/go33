package main

import "fmt"

type Flight struct {
	price int
}

func BookFlights(flights ...Flight) int {
	totalCost := 0

	for _, flight := range flights {
		totalCost += flight.price
	}
	return totalCost
}

func main() {
	flight1 := Flight{price: 100}
	flight2 := Flight{price: 200}

	fmt.Println(BookFlights(flight1, flight2))

	flight3 := Flight{price: 100}
	flight4 := Flight{price: 250}
	flight5 := Flight{price: 300}
	fmt.Println(BookFlights(flight1, flight2, flight3, flight4, flight5))
}
