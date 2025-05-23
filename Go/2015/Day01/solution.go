package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()

	timesGoUp := strings.Count(input, "(")
	timesGoDown := strings.Count(input, ")")

	finalFloor := timesGoUp - timesGoDown

	fmt.Printf("Santa ends up on: %d\n", finalFloor)
}

func readInput() string {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Unable to read file")
		return "ERROR"
	}

	return string(f)
}
