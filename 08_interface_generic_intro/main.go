package main

import "fmt"

func main() {
	result := add("5", "10")
	fmt.Println(result)
}

// func add(a, b interface{}) interface{} {
// 	aInt, aIsInt := a.(int)
// 	bInt, bIsInt := b.(int)

// 	if aIsInt && bIsInt {
// 		return aInt + bInt
// 	}

// 	aFloat, aIsFloat := a.(float64)
// 	bFloat, bIsFloat := b.(float64)
// 	if aIsFloat && bIsFloat {
// 		return aFloat + bFloat
// 	}
// 	aStr, aIsStr := a.(string)
// 	bStr, bIsStr := b.(string)
// 	if aIsStr && bIsStr {
// 		return aStr + bStr
// 	}
// 	return nil
// }

func add[T int | float64 | string](a, b T) T {
	return a + b
}
