package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	off = iota
	on
	toggle
)

type Lights struct {
	light map[string]bool
}

func NewLightsMap() *Lights {
	return &Lights{
		light: make(map[string]bool),
	}
}

func (lm *Lights) SwitchLights(action int, startCoords, endCoords []int) {
	x_start := startCoords[0]
	x_end := endCoords[0]
	y_start := startCoords[1]
	y_end := endCoords[1]

	for y := y_start; y <= y_end; y++ {
		for x := x_start; x <= x_end; x++ {
			keyString := fmt.Sprintf("(%d,%d)", x, y)
			switch action {
			case 0:
				lm.light[keyString] = false
			case 1:
				lm.light[keyString] = true
			case 2:
				_, ok := lm.light[keyString]
				if ok {
					lm.light[keyString] = !lm.light[keyString]
				} else {
					lm.light[keyString] = true
				}
			}
		}
	}
}

func (lm *Lights) litLights() {
	litLightsCount := 0

	for _, val := range lm.light {
		if val {
			litLightsCount++
		}
	}
	fmt.Printf("Number of lit lights: %d\n", litLightsCount)
}

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	flag.Parse()

	input, err := readInput(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	xmas_lights := NewLightsMap()

	for _, ti := range input {
		xmas_lights.SwitchLights(parseInstructions(ti))
	}
	xmas_lights.litLights()
}

func parseInstructions(instruction string) (int, []int, []int) {
	startCoordinates, endCoordinates := getCoordinates(instruction)
	action := getAction(instruction)
	return action, startCoordinates, endCoordinates
}

func getAction(instruction string) int {
	instructionTypes := []string{"turn on", "turn off"}

	switch {
	case strings.Contains(instruction, instructionTypes[0]):
		return on
	case strings.Contains(instruction, instructionTypes[1]):
		return off
	default:
		return toggle
	}
}

func getCoordinates(instruction string) ([]int, []int) {
	instructionSlice := strings.Split(instruction, " ")

	startCoordinateString := strings.Split(instructionSlice[len(instructionSlice)-3], ",")
	endCoordinateString := strings.Split(instructionSlice[len(instructionSlice)-1], ",")

	return convertCoordinatesToInt(startCoordinateString), convertCoordinatesToInt(endCoordinateString)
}

func convertCoordinatesToInt(coords []string) []int {
	intCoords := make([]int, 0, 2)
	for _, c := range coords {
		intCoord, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("Error converting string to int")
			os.Exit(1)
		}
		intCoords = append(intCoords, intCoord)
	}
	return intCoords
}

// Read Input

func readInput(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file '%s': %w", filename, err)
	}

	content := strings.Split(string(data), "\n")
	return content, nil
}
