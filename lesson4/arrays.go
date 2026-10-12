package main

import "fmt"

func main() {
	var stockPrices [5]float64
	stockPrices = [5]float64{101.2, 98.4, 70.1, 110.4, 154.4}
	fmt.Println(stockPrices)

	userProfiles := [...]struct {
		Name  string
		Email string
	}{
		{"FirstName1 LastName1", "firstname1.lastname1@example.com"},
		{"FirstName2 LastName2", "firstname2.lastname2@example.com"},
	}
	fmt.Println(userProfiles)

	stockSymbols := [...]string{"APPL", "GOOGL", "AMZN"}
	fmt.Println(stockSymbols)

	stocks := [3]string{"Apple", "Google", "Amazon"}
	fmt.Println(stocks[0])

	var userAges [5]int

	userAges[0] = 25
	userAges[2] = 30

	fmt.Println(userAges)

	stockPrices1 := [3]float64{142.2, 150.1, 137.2}
	stockPrices1[2] = 72.1
	fmt.Println(stockPrices1)

}
