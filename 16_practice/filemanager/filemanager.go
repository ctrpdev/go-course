package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath   string
	OutpuhtFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	defer file.Close() // defer llama a file.Close() al final de la función, asegurando que el archivo se cierre incluso si ocurre un error antes de llegar a esa línea.

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, fmt.Errorf("error while reading file: %w", err)
	}

	// file.Close()

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutpuhtFilePath)

	if err != nil {
		return errors.New("Failed to create file: " + err.Error())
	}
	defer file.Close()
	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Failed to write JSON: " + err.Error())
	}

	// file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:   inputPath,
		OutpuhtFilePath: outputPath,
	}
}
