package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()

	// timesGoUp := strings.Count(input, "(")
	// timesGoDown := strings.Count(input, ")")

	// finalFloor := timesGoUp - timesGoDown
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

func readInput() string {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(1)
	}

	return string(f)
}
