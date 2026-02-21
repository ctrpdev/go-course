package main

import (
	"fmt"
)

// func main() {
func ex004() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit

	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Earnings after tax: ", profit)
	fmt.Println("Ratio: ", ratio)
}
