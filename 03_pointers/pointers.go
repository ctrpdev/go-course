package main

import "fmt"

func ex001() {
	// func main() {
	age := 32 // regular variable

	var agePointer *int

	agePointer = &age // pointer variable

	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", *agePointer)

	// adultYears := getAdultYears(&age)
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
