package main

import "fmt"

func main() {
	str := "Hello"

	for i := 0; i < len(str); i++ {
		fmt.Printf("Character at index %d is: %c\n", i, str[i])
	}

	for i, v := range str {
		fmt.Printf("Character at index %d is: %c\n", i, v)
	}
}
