package main

import "fmt"

// func main() {
func ex009() {
	accountBalance := 1000.0

	fmt.Println("Welcome to GO Bank")

	// for range 2 {}
	// for i := 0; i < 2; i++ {
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		if choice == 1 {
			fmt.Println("Your balance is:", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("New balance is:", accountBalance)
		} else if choice == 3 {
			var WithdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&WithdrawAmount)

			if WithdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			if WithdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Must be bigger or equal than your balance.")
				// return
				continue
			}

			accountBalance -= WithdrawAmount
			fmt.Println("New balance is:", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			// return
			break
		}
	}

	fmt.Println("Thank you for using GO Bank!")
}
