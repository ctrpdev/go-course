package main

import "fmt"

type str string

func (s str) log() {
	fmt.Println(s)
}

func main() {
	var name str = "Nombre custom"
	name.log()
}
