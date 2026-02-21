package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// func ex011() {
	revenue, err := getInputs("Enter revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getInputs("Enter expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getInputs("Enter tax rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	ebt, profit, ratio := getCalculation(revenue, expenses, taxRate)

	writeToFile(ebt, profit, ratio)
	formattedEBT, formattedProfit, formattedRatio := formatValue(ebt, profit, ratio)
	fmt.Print(formattedEBT, formattedProfit, formattedRatio)
}

func getInputs(label string) (float64, error) {
	var input float64
	var err error
	fmt.Print(label)
	fmt.Scan(&input)

	if input < 1 {
		err = errors.New("input cannot be zero nor negative")
		return 0.0, err
	}

	return input, nil
}

func getCalculation(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func formatValue(ebt, profit, ratio float64) (string, string, string) {
	formattedEBT := fmt.Sprintf("Earnings before tax: %.1f\n", ebt)
	formattedProfit := fmt.Sprintf("Earnings after tax: %.1f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.1f\n", ratio)
	return formattedEBT, formattedProfit, formattedRatio
}

func writeToFile(ebt float64, profit float64, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f", ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(results), 0644)
}
