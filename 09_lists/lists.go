package main

import "fmt"

// func main() {
func main() {
	prices := []float64{10.99, 8.99} // crea un slice
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	fmt.Println(prices)
	updatedPrices := append(prices, 12.99) // append crea un nuevo slice
	fmt.Println(updatedPrices)             // por eso aca estan todos los nuevos valores
	fmt.Println(prices)                    // pero aca sigue igual, porque append no modifica el slice original

	// para modificar el slice original, hay que reasignarlo
	prices = append(prices, 12.99)
	fmt.Println(prices)
	// debido a que un array tiene un tamaño fijo, los slices son más flexibles y se usan con más frecuencia en Go.
	// para eliminar un elemento de un slice, se puede usar append junto con slicing
	// eliminar el primer elemento
	prices = prices[1:] // crea un nuevo slice sin el primer elemento
	fmt.Println(prices)

	// unpacking lists
	discountPrices := []float64{10.99, 5.99, 20.00}
	prices = append(prices, discountPrices...) // los tres puntos indican que se deben agregar todos los elementos del slice
	fmt.Println(prices)

}

// for reference
// func main() {
// 	// var productNames [4]string
// 	// productNames = [4]string{"A Book"}	// crea un array de 4 elementos
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Guitar"

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	priceSlice := prices[1:3]
// 	fmt.Println(priceSlice)

// 	fmt.Println(productNames[2])

// 	// empezar desde el principio
// 	slicePrices2 := prices[:3]
// 	fmt.Println(slicePrices2)

// 	// hasta el final
// 	slicePrices3 := prices[2:]
// 	fmt.Println(slicePrices3)

// 	// no se puede usar -1 para empezar desde el final
// 	// slicePrices4 := prices[1:4]
// 	// fmt.Println(slicePrices4)

// 	// no se puede usar un index mayor al tamaño del array, solo se puede el index + 1 del tamaño del array
// 	// slicePrices5 := prices[1:5]
// 	// fmt.Println(slicePrices5)

// 	// se puede usar un slice dentro de otro
// 	featuredPrices := prices[1:]
// 	hightlitedPrices := featuredPrices[:1]
// 	fmt.Println(hightlitedPrices)

// 	// cambiar un valor en el slice cambia el array original
// 	featuredPrices[0] = 1000.0
// 	fmt.Println(prices)

// 	// eso es porque el slice es una referencia al array original
// 	// al crear un slice no se copia el array, solo se crea una referencia al array original

// 	// obtener el largo del array y la capacidad del slice
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))

// 	// len es el numero de elementos en el array, es la capacidad del array, y como hightlitedPrices es un slice de featuredPrices y featuredPrices tiene un cap de 3, hightlitedPrices tambien tiene un cap de 3
// 	fmt.Println(len(hightlitedPrices), cap(hightlitedPrices))

// 	// siempre se puede seleccionar mas hacia adelante pero no hacia atras
// 	hightlitedPrices = featuredPrices[:3]
// 	fmt.Println(hightlitedPrices)
// 	fmt.Println(len(hightlitedPrices), cap(hightlitedPrices))

// }
