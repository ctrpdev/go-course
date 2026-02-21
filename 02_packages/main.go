package main

import (
	"fmt"

	"example.com/ctrp/packages/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile, 0.0)

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	fmt.Println("Welcome to GO Bank")
	fmt.Println("Phone number: ", randomdata.PhoneNumber())

	for {
		showMenu()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("New balance is:", accountBalance)
		case 3:
			var WithdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&WithdrawAmount)

			if WithdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if WithdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Must be bigger or equal than your balance.")
				continue
			}

			accountBalance -= WithdrawAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("New balance is:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using GO Bank!")
			return
		}
	}

}
