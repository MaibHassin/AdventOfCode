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

	coordinates := make(map[string]bool, len(input))

	x, y := 0, 0
	coordinates[fmt.Sprintf("%d,%d", x, y)] = true
	for _, d := range strings.Split(input, "") {
		switch d {
		case "^":
			y += 1
			coordinates[fmt.Sprintf("%d,%d", x, y)] = true
		case "v":
			y -= 1
			coordinates[fmt.Sprintf("%d,%d", x, y)] = true
		case ">":
			x += 1
			coordinates[fmt.Sprintf("%d,%d", x, y)] = true
		case "<":
			x -= 1
			coordinates[fmt.Sprintf("%d,%d", x, y)] = true
		}

	}
	fmt.Println(len(coordinates))
}

func readInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file '%s': %w", filename, err)
	}

	content := strings.TrimSpace(string(data))
	return content, nil
}
