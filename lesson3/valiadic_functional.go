package main

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
}
