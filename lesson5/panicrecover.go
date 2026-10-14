package main

import (
	"fmt"
	"time"
)

func bookFlight() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from: ", r)
		}
	}()

	fmt.Println("Attempting to book a flight....")
	time.Sleep(2 * time.Second)
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error while booking: ", r)
		}
	}()

	fmt.Println("Attempting to book a flight....")

	panic("Fail to book the flight due to system error.")
}
