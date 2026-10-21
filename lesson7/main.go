package main

import (
	"fmt"

	"lesson7/calculator"

	"github.com/nsf/termbox-go"
)

func main() {

	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer termbox.Close()

	result := calculator.Add(5, 3)
	fmt.Println(result)
}
