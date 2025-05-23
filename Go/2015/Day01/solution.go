package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	flag.Parse()

	input, err := readInput(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	finalFloor := 0

	firstNeagtiveFound := false
	var firstNegative int

	for i, floor := range strings.Split(input, "") {
		if floor == "(" {
			finalFloor++
		} else {
			finalFloor--
		}

		if finalFloor < 0 && !firstNeagtiveFound {
			firstNegative = i + 1
			firstNeagtiveFound = true
		}
	}

	fmt.Printf("Santa ends up on: %d\n", finalFloor)
	fmt.Printf("First position when santa enters the basement: %d\n", firstNegative)
}

func readInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file '%s': %w", filename, err)
	}
	content := strings.TrimSpace(string(data))
	return content, nil
}
