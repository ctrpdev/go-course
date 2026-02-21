package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 0.0, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0.0, errors.New("Failed to parse balance.")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

// func main() {
func ex010() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		// return // si no se quiere continuar con la ejecucion
		panic(err) // si no se quiere continuar con la ejecucion, indicando donde se genero el panic
	}

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

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("New balance is:", accountBalance)
		case 3:
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
			writeBalanceToFile(accountBalance)
			fmt.Println("New balance is:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using GO Bank!")
			return
		}
	}

}
