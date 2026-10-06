package main

import "errors"

func percentChange(oldPrice float64, newPrice float64) (float64, error) {

	if oldPrice == 0.0 {
		return 0, errors.New("old value cannot be zero")
	}

	return ((newPrice - oldPrice) / oldPrice) * 100, nil
}
