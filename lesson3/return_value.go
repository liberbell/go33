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

func greet(name string) string {
	return "Hello, " + name + "!"
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	change, err := percentChange(10000.0, 120.0)

	if err != nil {
		fmt.Println("An error occurred:", err.Error())
		return
	}

	fmt.Println("The percentage change is:", change)

	quotient, remainder := divide(10, 3)
	fmt.Printf("The Quotient is %d, and the Rreemaindr is %d\n", quotient, remainder)

	fmt.Println(greet("James"))
	a, b := 5, 10
	a, b = swap(a, b)
}
