package main

import "fmt"

func main() {
	var age int = 20
	fmt.Println("Age:", age)

	var pi float64 = 3.14159
	fmt.Println("Value of Pi:", pi)

	if score := 85; score >= 50 {
		fmt.Println("Passed with score:", score)
	} else {
		fmt.Println("Failed with score:", score)
	}
}
