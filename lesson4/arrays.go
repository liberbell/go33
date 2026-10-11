package main

import "fmt"

func main() {
	var stockPrices [5]float64
	stockPrices = [5]float64{101.2, 98.4, 70.1, 110.4, 154.4}
	fmt.Println(stockPrices)

	userProfiles := [...]struct {
		Name  string
		Email string
	}(
		{"FirstName1 LastName1", "firstname1.lastname1@example.com"},
		{"FirstName2 LastName2", "firstname2.lastname2@example.com"},
	)

}
