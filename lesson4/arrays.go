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
	for _, symbol := range stockSymbols {
		fmt.Println("Symbol is ", symbol)
	}

	var userAges [5]int

	userAges[0] = 25
	userAges[2] = 30

	fmt.Println(userAges)

	stockPrices1 := [3]float64{142.2, 150.1, 137.2}
	stockPrices1[2] = 72.1
	fmt.Println(stockPrices1)

	users := [3]string{"James", "money", "Q"}
	totalUsers := len(users)
	fmt.Println("Total number of users:", totalUsers)

	userProfiles := [2][2]struct {
		Name string
		Email string
	}{
		{{"Firstname1 Lastname1", "firstname1.lastname1@example.com"}, {"Firstname2 Lastname2", "firstname2.lastname2@example.com"}}
		{{"Firstname3 Lastname3", "firstname3.lastname3@example.com"}, {"Firstname4 Lastname4", "firstname4.lastname4@example.com"}}
	}
}
