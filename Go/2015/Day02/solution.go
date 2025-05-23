package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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

	parsedDimensions := parseInput(input)
	fmt.Println(calculateRequiredPaper(parsedDimensions))
	fmt.Println(calculateRequiredRibbon(parsedDimensions))
}

func parseInput(input string) [][]string {
	dimenstionsList := strings.Split(input, "\n")
	var splitDimensions [][]string
	for _, gd := range dimenstionsList {
		splitDimensions = append(splitDimensions, strings.Split(gd, "x"))
	}
	return splitDimensions
}

func calculateRequiredPaper(list [][]string) int {
	totalPaper := 0
	for _, d := range list {
		l, _ := strconv.Atoi(d[0])
		w, _ := strconv.Atoi(d[1])
		h, _ := strconv.Atoi(d[2])
		lw := l * w
		wh := w * h
		hl := h * l
		smallestSide := min(lw, wh, hl)
		requiredPaper := (2 * lw) + (2 * wh) + (2 * hl) + smallestSide
		totalPaper += requiredPaper
	}
	return totalPaper
}

func calculateRequiredRibbon(list [][]string) int {
	totalRibbon := 0

	for _, d := range list {
		l, _ := strconv.Atoi(d[0])
		w, _ := strconv.Atoi(d[1])
		h, _ := strconv.Atoi(d[2])
		ribbonPresent := (l * 2) + (w * 2) + (h * 2) - (max(l, w, h) * 2)
		ribbonBow := l * w * h
		totalRibbon += ribbonPresent + ribbonBow
	}
	return totalRibbon
}

func readInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file '%s': %w", filename, err)
	}

	content := strings.TrimSpace(string(data))
	return content, nil
}
