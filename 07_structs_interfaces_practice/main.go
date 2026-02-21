package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// si tengo una interface con un solo metodo, el nombre de la inrterface por convencion el nombre del metodo + "er"
type saver interface {
	Save() error
}

// type getter interface {
// 	Get()
// }

// otra forma seria definir la interface completa
// type outputtable interface {
// 	Save() error
// 	Get()
// }

// o definir la interface compuesta
// type outputtable interface {
// 	saver
// 	getter
// }

// o definir la interface compuesta directamente
type outputtable interface {
	saver
	Get()
}

func main() {
	printSomething(1)
	printSomething(1.4)
	printSomething("hola")
	title, content := getNoteData()
	todoText := getUserInput("Todo:")
	todoItem, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	err = outputData(todoItem)
	if err != nil {
		return
	}

	outputData(userNote)
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}
	fmt.Println("Data saved successfully.")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func outputData(data outputtable) error {
	data.Get()
	return saveData(data)
}

// func printSomething(something interface{}) {
// 	fmt.Println(something)
// }

func printSomething(value any) {
	switch value.(type) {
	case int:
		fmt.Println("Es un entero:", value)
	case float64:
		fmt.Println("Es un float64:", value)
	case string:
		fmt.Println("Es una cadena:", value)
		// default:
		// 	fmt.Println("Tipo desconocido")
	}
}
