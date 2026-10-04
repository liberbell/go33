package main

import "fmt"

func main() {
	num := 10

	if num%2 == 0 {
		fmt.Println(num, "is even")
	}

	X := 100
	if X == 100 {
		fmt.Println("Japan")
	} else if X == 200 {
		fmt.Println("China")
	} else {
		fmt.Println("Germany")
	}
}
