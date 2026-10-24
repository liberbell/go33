package main

import (
	simpleGoApp "SimpleGoApp"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askForNumber() int {
	fmt.Println("Please enter a number to add to the data structure:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	number, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return askForNumber
	}

	return number
}

func askForConfirmation(message string) bool {
	fmt.Println("message " + "(yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = response[:len(response)-1]
	return response == "yes"
}

func main() {
	data := simpleGoApp.NewData()

	number := askForNumber()

	data.AddElement(number)
	fmt.Printf("Number %d added to the data structure.\n", number)
}
