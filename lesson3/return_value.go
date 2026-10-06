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

func divide(a int, b int) (quo int, rem int) {
	quo = a / b
	rem = a % b

	return
}

func main() {
	change, err := percentChange(10000.0, 120.0)

	if err != nil {
		fmt.Println("An error occurred:", err.Error())
		return
	}

	fmt.Println("The percentage change is:", change)
}
