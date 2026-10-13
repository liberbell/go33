package main

import "fmt"

func main() {
	userProfiles := map[string]string{
		"FirstName_1": "Name_1",
		"Email1":      "Email_1@examle.com",
		"Phone_1":     "123-456-7890",
	}

	fmt.Println(userProfiles)
	fmt.Println("Firstname: ", userProfiles["FirstName_1"])
	fmt.Println("Email: ", userProfiles["Email1"])

	userProfile2 := map[string]string{
		"FirstName_2": "Name_2",
		"Email2":      "Email_2@example.com",
		"Phone2":      "111-222-3333",
	}
	userProfile2["Email2"] = "Email_21@example.com"
}
