package main

import (
	"fmt"
	"math"
)

// func main() {
func ex002() {
	const inflationRate = 2.5
	var years float64 = 10
	var investmentAmount float64 = 1000
	expectedReturnRate := 0.05

	// Different ways to declare variables
	// var investmentAmount, years = 1000, "10"
	// var investmentAmount, years float64 = 1000, 10
	// investmentAmount, years := 1000.0, 10.0

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value: ", futureValue)
	fmt.Println("Future real value: ", futureRealValue)
}
