package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Input cannot be empty")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Get() {
	fmt.Print(todo.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		return err
	}
	return nil
}
