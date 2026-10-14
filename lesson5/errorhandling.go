package main

import "errors"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero is not allowed")
	}
	return a / b, nil
}
