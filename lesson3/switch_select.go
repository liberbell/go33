package main

import "fmt"

func main() {
	accountStatus := "active"

	switch accountStatus {
	case "Active":
		fmt.Println("Your account is active.")
		fallthrough
	case "Inactive":
		fmt.Println("Your account is inactive.")
	default:
		fmt.Println("Your account status is unknown")
	}

}
