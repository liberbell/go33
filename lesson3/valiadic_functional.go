package main

type Flight struct {
	price int
}

func BookFlights(flights ...Flight) int {
	totalCost := 0

	for _, flight := range flights {
		totalCost += flight.price
	}
}
