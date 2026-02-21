package main

import (
	"fmt"
)

// func main() {
func ex008() {
	revenue := getInput("Enter revenue: ")
	expenses := getInput("Enter expenses: ")
	taxRate := getInput("Enter tax rate: ")

	ebt, profit, ratio := getCalculations(revenue, expenses, taxRate)

	formattedEBT, formattedProfit, formattedRatio := formatValues(ebt, profit, ratio)
	fmt.Print(formattedEBT, formattedProfit, formattedRatio)
}

func getInput(label string) float64 {
	var input float64
	fmt.Print(label)
	fmt.Scan(&input)
	return input
}

func getCalculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func formatValues(ebt, profit, ratio float64) (string, string, string) {
	formattedEBT := fmt.Sprintf("Earnings before tax: %.1f\n", ebt)
	formattedProfit := fmt.Sprintf("Earnings after tax: %.1f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.1f\n", ratio)
	return formattedEBT, formattedProfit, formattedRatio
}
