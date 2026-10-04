package main

import "fmt"

func main() {
	var message1 = "Hello world"
	message2 := "Hello golang"
	fmt.Println(message1)
	fmt.Println(message2)

	name := "Programmer"
	fmt.Println("%c\n", name[0])
	fmt.Println("%c\n", name[3])
	fmt.Println("%c\n", name[8])
}
