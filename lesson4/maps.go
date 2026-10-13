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
	fmt.Println("UserProfile2: ", userProfile2)

	userProfile3 := map[string]string{
		"FirstName_3": "Name_3",
		"Email3":      "Email_3@example.com",
		"Phone3":      "111-222-3333",
	}
	userProfile3["LastName3"] = "LastName_3"
	fmt.Println(userProfile3)

	delete(userProfile3, "Phone3")
	fmt.Println("deleted userprofile3: ", userProfile3)

	userProfile5 := map[string]string{
		"FirstName_5": "Name_5",
		"Email_5":     "Email_5@example.com",
		"Phone_5":     "111-222-3333",
	}
	for key, value := range userProfile5 {
		fmt.Printf("%s: %s\n", key, value)
	}

	bankAccount := make(map[string]float64, 0)
	bankAccount["Name_1"] = 1000.50
	fmt.Println(bankAccount)
}
