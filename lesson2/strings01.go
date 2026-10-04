package main

import "fmt"

func main() {
	var message1 = "Hello world"
	message2 := "Hello golang"
	fmt.Println(message1)
	fmt.Println(message2)

	name := "Programmer"
	fmt.Printf("%c\n", name[0])
	fmt.Printf("%c\n", name[3])
	fmt.Printf("%c\n", name[8])

	message := "welcome to Programmer world"
	stringLength := len(message)
	println("Length of a string is:", stringLength)
}
