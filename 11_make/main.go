package main

import "fmt"

// tambien se puede crear un alias para un tipado de dato largo, a fin de tipar menos, pero es a gusto del programador
type strIntMap map[string]int

// incluso se pueden definir metodos para estos alias
func (m strIntMap) outputMap() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) // crea un slice de strings con longitud 2 y capacidad 5, para evitar realocaciones de memoria
	userNames[0] = "Alice"
	userNames[1] = "Bob"

	fmt.Println("User Names:", userNames)

	userNames = append(userNames, "Charlie", "Diana") // agrega más elementos al slice
	fmt.Println("Updated User Names:", userNames)

	// tambien funciona con mapas
	userAges := make(map[string]int, 2) // crea un mapa con capacidad inicial de 2 pares clave-valor
	userAges["Alice"] = 30
	userAges["Bob"] = 25

	fmt.Println("User Ages:", userAges)

	userAges["Charlie"] = 28 // agrega más pares clave-valor al mapa
	fmt.Println("Updated User Ages:", userAges)

	newUserAres := make(strIntMap, 2) // usando el alias definido arriba
	newUserAres["Diana"] = 22
	newUserAres["Eve"] = 35
	newUserAres.outputMap() // usando el metodo definido para el alias

	// FOR
	for i, v := range userNames {
		fmt.Println("Index: ", i, "| Value: ", v)
	}

	for k, v := range userAges {
		fmt.Println("Key:", k, "| Value:", v)
	}
}
