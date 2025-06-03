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
	grid [1000][1000]bool
}

func NewLightsGrid() *Lights {
	return &Lights{}
}

func (lm *Lights) SwitchLights(action int, startCoords, endCoords []int) {
	x_start := startCoords[0]
	x_end := endCoords[0]
	y_start := startCoords[1]
	y_end := endCoords[1]

	for y := y_start; y <= y_end; y++ {
		for x := x_start; x <= x_end; x++ {

			switch action {
			case off:
				lm.grid[x][y] = false
			case on:
				lm.grid[x][y] = true
			case toggle:
				lm.grid[x][y] = !lm.grid[x][y]
			}
		}
	}
}

func (lm *Lights) litLights() {
	litLightsCount := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if lm.grid[x][y] {
				litLightsCount++
			}
		}
	}
	fmt.Printf("Number of lit lights: %d\n", litLightsCount)
}

type LightsWithBrightness struct {
	grid [1000][1000]int
}

func NewLightsWithBrightnessMap() *LightsWithBrightness {
	return &LightsWithBrightness{}
}

func (lm *LightsWithBrightness) changeBrightness(action int, startCoords, endCoords []int) {
	x_start := startCoords[0]
	x_end := endCoords[0]
	y_start := startCoords[1]
	y_end := endCoords[1]

	for y := y_start; y <= y_end; y++ {
		for x := x_start; x <= x_end; x++ {

			switch action {
			case off:
				if lm.grid[x][y]-1 < 0 {
					lm.grid[x][y] = 0
				} else {
					lm.grid[x][y] -= 1
				}
			case on:
				lm.grid[x][y] += 1
			case toggle:
				lm.grid[x][y] += 2
			}
		}
	}
}

func (lm *LightsWithBrightness) litLightsBrightness() {
	totalBrightness := 0
	for _, row := range lm.grid {
		for _, value := range row {
			totalBrightness += value
		}
	}

	fmt.Printf("The total brightness of lights is: %d\n", totalBrightness)
}

func main() {
	inputFileName := flag.String("input", "../input.txt", "Path to the input file")
	part := flag.Int("part", 1, "Please select which part do you want to run (1 or 2)")
	flag.Parse()

	input, err := readInput(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *part {
	case 1:
		Part1(input)
	case 2:
		Part2(input)
	}

}

func Part1(input []string) {
	xmasLights := NewLightsGrid()

	for _, ti := range input {
		xmasLights.SwitchLights(parseInstructions(ti))
	}
	xmasLights.litLights()
}

func Part2(input []string) {
	lightsBrightness := NewLightsWithBrightnessMap()

	for _, ti := range input {
		lightsBrightness.changeBrightness(parseInstructions(ti))
	}
	lightsBrightness.litLightsBrightness()
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
