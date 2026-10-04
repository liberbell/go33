package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "admin 01"
	string2 := "admin pro"
	string3 := "admin 03"
	fmt.Println(strings.Compare(string1, string2))
	fmt.Println(strings.Compare(string1, string3))
	fmt.Println(strings.Compare(string2, string3))
}
