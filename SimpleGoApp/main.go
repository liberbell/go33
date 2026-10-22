package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askForNumber() int {
	fmt.Println("Please enter a number to add to the data structure:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	number, err := strconv.Atoi()
}
