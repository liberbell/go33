package main

import "fmt"

func main() {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println("Size of KB:", KB)
	fmt.Println("Size of MB:", MB)
	fmt.Println("Size of GB:", GB)
	fmt.Println("Size of TB:", TB)
}
