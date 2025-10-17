package main

import (
	"fmt"
	"main/user"
)

func main() {
	// Version 1
	// firstName := getUserData("Please enter your first name: ")
	// lastName := getUserData("Please enter your last name: ")
	// birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// userData := user{
	// 	FirstName: firstName,
	// 	LastName: lastName,
	// 	Birthdate: birthdate,
	// 	CreatedAt: time.Now().UTC().Format(time.RFC3339),
	// 	// Age: calculateAge(birthdate),
	// }

	// Version 2
	// userData := user{
	// 	FirstName: getUserData("Please enter your first name: "),
	// 	LastName: getUserData("Please enter your last name: "),
	// 	Birthdate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
	// 	CreatedAt: time.Now().UTC().Format(time.RFC3339),
	// }

	// userData.outputUserData()

	// Version 3
	userData, err := user.New()
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println(userData)

	adminData, err := user.ChangeUserToAdmin(userData)
	if err != nil {
		fmt.Println("Error changing user to admin:", err)
		return
	}
	fmt.Println(adminData)
}

