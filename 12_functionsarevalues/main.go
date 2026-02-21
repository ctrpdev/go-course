package main

import "fmt"

// para acordar el texto en los parametros se puede tipear la funcion que se espera
type transformFunc func(int) int

// ejemplo
// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	doubled := transformNumbers(&numbers, double) // si se usa una funcion como parametro no se usan los ()
	fmt.Println(doubled)                          // Output: [2 4 6 8 10]

	tripled := transformNumbers(&numbers, tiple)
	fmt.Println(tripled) // Output: [3 6 9 12 15]

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	fmt.Println(transformNumbers(&numbers, transformerFn1))
	fmt.Println(transformNumbers(&moreNumbers, transformerFn2))
}

// func transformNumbers(numbers *[]int, transform func(int) int) []int {
func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return tiple // cuando se retorna una funcion no se usan los ()
	}
}

func double(n int) int {
	return n * 2
}

func tiple(n int) int {
	return n * 3
}
