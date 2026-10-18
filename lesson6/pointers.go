package main

import "fmt"

func main() {

	var n = 10
	var p = &n
	fmt.Printf("Type of Pointer: %T\n", p)

	fmt.Printf("Address: %v\n", p)

	fmt.Printf("Value: %v\n", *p)

	var a int = 50
	var b *int = &a

	fmt.Println(a)
	fmt.Println(*b)
}
