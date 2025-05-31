package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	numAgents := flag.Int("agents", 1, "Number of agents")
	flag.Parse()

	input, err := readInput(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	minHousesVisited := countHousesVisited(input, *numAgents)
	fmt.Printf("Total unique houses visited by %d agent(s) are %d\n", *numAgents, minHousesVisited)

}

type Point struct {
	X int
	Y int
}

func countHousesVisited(c string, numOfAgents int) int {
	coordinates := make(map[Point]bool, len(c))
	agents := make([]Point, numOfAgents)

	for _, initialPos := range agents {
		coordinates[initialPos] = true
	}
	for i, d := range c {
		agentNumber := i % numOfAgents

		agents[agentNumber] = move(d, agents[agentNumber])
		coordinates[agents[agentNumber]] = true
	}
	return len(coordinates)
}

func move(d rune, currentPos Point) Point {
	newPos := currentPos
	switch d {
	case '^':
		newPos.Y += 1
	case 'v':
		newPos.Y -= 1
	case '>':
		newPos.X += 1
	case '<':
		newPos.X -= 1
	}
	return newPos
}

func readInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file '%s': %w", filename, err)
	}

	content := strings.TrimSpace(string(data))
	return content, nil
}
