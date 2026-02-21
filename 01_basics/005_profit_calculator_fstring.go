package main

import (
	"fmt"
)

// func main() {
func ex005() {
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

	formattedEBT := fmt.Sprintf("Earnings before tax: %.1f\n", ebt)
	formattedProfit := fmt.Sprintf("Earnings after tax: %.1f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.1f\n", ratio)
	// fmt.Printf("Earnings before tax: %v\nEarnings after tax: %v\nRatio: %v\n", ebt, profit, ratio)
	// fmt.Printf("Earnings before tax: %.1f\nEarnings after tax: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	fmt.Print(formattedEBT, formattedProfit, formattedRatio)
	// fmt.Printf(`
	// Earnings before tax: %v
	// Earnings after tax: %v
	// Ratio: %v`,
	// 	ebt,
	// 	profit,
	// 	ratio)

}
