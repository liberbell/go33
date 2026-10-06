package main

import (
	"errors"
	"fmt"
)

func percentChange(oldPrice float64, newPrice float64) (float64, error) {

	if oldPrice == 0.0 {
		return 0, errors.New("old value cannot be zero")
	}

	return ((newPrice - oldPrice) / oldPrice) * 100, nil
}

func main() {
	change, err := percentChange(100.0, 120.0)

	if err != nil {
		fmt.Println("An error occurred:", err.Error())
		return
	}

	fmt.Println("The percentage change is:", change)
}
