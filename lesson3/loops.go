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

	m := map[string]string{
		"one":   "first",
		"two":   "second",
		"three": "third",
	}
	for k, v := range map {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
}
