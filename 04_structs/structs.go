package main

import (
	"fmt"

	"example.com/ctrp/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	// Un ejemplo de como el Case aplica a los atributos y metodos de un struct
	// appUser = &user.User{
	// 	FirstName: firstName,
	// }

	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	admin := user.NewAdmin("test", "pass")

	admin.OutputDetails()
	admin.Clear()
	admin.OutputDetails()

	appUser.OutputDetails()
	appUser.Clear()
	appUser.OutputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
