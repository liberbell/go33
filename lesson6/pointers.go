package main

import "fmt"

func main() {

	var n = 10
	var p = &n
	fmt.Printf("Type of Pointer: %T\n", p)

	fmt.Printf("Address: %v\n", p)
}
