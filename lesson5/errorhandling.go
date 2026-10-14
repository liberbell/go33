package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	result, err := divide(4, 2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %f\n", result)

	_, err = divide(4, 0)

	if err != nil {
		log.Printf("Error: %s\n", err)
	}
}
