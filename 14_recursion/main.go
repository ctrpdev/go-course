package main

import "fmt"

func main() {
	fact := noRecursiveFactorial(5)
	fmt.Println(fact)

	fact2 := recursiveFactorial(5)
	fmt.Println(fact2)
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120
func noRecursiveFactorial(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result
}

func recursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}
