package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	// unpacking slice
	sum := sumup(1, 10, 15, 40, -5)
	// param... unpacks the slice into individual arguments
	anotherSum := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// se puede aceptar cualquier tipo de valor como un ...type y los parametros fijos van antes
// func sumup(startingValue int, nums ...int) int {
// ...type collects all the arguments into a slice of ints
func sumup(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
