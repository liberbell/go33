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

	email := "apple@bank.com"
	containBank := strings.Contains(email, "banks")
	fmt.Println(containBank)

	oldEmail := "apple@citibank.com"
	newEmail := strings.Replace(oldEmail, "citi", "chase", 1)
	fmt.Println(newEmail)

	profileName := "bankUserName"
	fmt.Println(strings.ToUpper(profileName))
	fmt.Println(strings.ToLower(profileName))
}
