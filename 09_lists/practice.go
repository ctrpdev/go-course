package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

// func main() {
func practice() {
	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	var hobbies [3]string = [3]string{"drawing", "playing guitar", "watching anime"}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	firstHobby := hobbies[0]
	fmt.Println(firstHobby)
	newHobbies := hobbies[1:3]
	fmt.Println(newHobbies)

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice1 := hobbies[0:2]
	fmt.Println(slice1)
	slice2 := hobbies[:2]
	fmt.Println(slice2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice3 := slice1[1:3]
	fmt.Println(slice3)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	dynamicArray := []string{"learn go", "create a go app"}
	fmt.Println(dynamicArray)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	dynamicArray[1] = "create a go backend"
	dynamicArray = append(dynamicArray, "record mayonaise")
	fmt.Println(dynamicArray)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		{
			id:    "1",
			title: "telecaster",
			price: 1399.99,
		},
		{
			id:    "2",
			title: "stratocaster",
			price: 1499.99,
		},
	}

	products = append(products,
		Product{
			id:    "3",
			title: "rg 550",
			price: 1199.99,
		},
	)

	fmt.Println(products)
}
