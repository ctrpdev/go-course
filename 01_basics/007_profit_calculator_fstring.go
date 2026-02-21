package main

import (
	"fmt"
)

// func main() {
func ex007() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue, expenses, taxRate = inputPrompt(revenue, expenses, taxRate)

	ebt, profit, ratio := calculation(revenue, expenses, taxRate)

	formattedEBT := fmt.Sprintf("Earnings before tax: %.1f\n", ebt)
	formattedProfit := fmt.Sprintf("Earnings after tax: %.1f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.1f\n", ratio)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)
}

func inputPrompt(revenue, expenses, taxRate float64) (r, e, t float64) {
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)
	return revenue, expenses, taxRate
}

func calculation(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
