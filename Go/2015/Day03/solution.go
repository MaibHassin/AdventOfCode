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
	xSanta, ySanta := 0, 0
	xRobot, yRobot := 0, 0

	for i, d := range strings.Split(input, "") {
		isEven := i%2 == 0
		if isEven {
			xSanta, ySanta = move(d, xSanta, ySanta)
			coordinates[fmt.Sprintf("%d,%d", xSanta, ySanta)] = true
		} else {
			xRobot, yRobot = move(d, xRobot, yRobot)
			coordinates[fmt.Sprintf("%d,%d", xRobot, yRobot)] = true
		}
	}
	fmt.Println(len(coordinates))
}

func move(d string, x, y int) (int, int) {
	switch d {
	case "^":
		y += 1
		return x, y
	case "v":
		y -= 1
		return x, y
	case ">":
		x += 1
		return x, y
	case "<":
		x -= 1
		return x, y
	}
	return 0, 0
}

func readInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file '%s': %w", filename, err)
	}

	content := strings.TrimSpace(string(data))
	return content, nil
}
