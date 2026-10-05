package main

import "fmt"

func square(num int) int {
	num *= num
	return num
}

func square2(ptr *int) {
	*ptr *= *ptr
}

func main() {
	original := 5

	result := square(original)
	fmt.Printf("The original value is %d and the square value is %d\n", original, result)

	square2(&original)
	fmt.Println("The original value after the square2 operation is %d\n", original)
}
