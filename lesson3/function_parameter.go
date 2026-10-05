package main

import "fmt"

func square(num int) int {
	num *= num
	return num
}

func main() {
	original := 5

	result := square(original)
	fmt.Printf("The original value is %d and the square value is %d\n", original, result)
}
