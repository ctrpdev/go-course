package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)

	getAgeDirectMutate(&age)
	fmt.Println(age)
}

func getAgeDirectMutate(age *int) {
	*age = *age - 18
}
