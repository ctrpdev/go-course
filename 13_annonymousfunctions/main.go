package main

import "fmt"

// una funcion anonima permite definir la funcion en el lugar donde se va a usar en vez de definirla por separado y luego pasarla como argumento

func main() {
	numbers := []int{1, 2, 3}

	doubled := createTransformer(2)
	tripled := createTransformer(3)

	fmt.Println(transformNumbers(&numbers, doubled))
	fmt.Println(transformNumbers(&numbers, tripled))

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}

	return dNumbers
}

// factory functions and closures (al usar un parametro dentro de un scope inferior, el parametro se "cierra" dentro de la funcion anonima)
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
