package main

import (
	"errors"
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
}
