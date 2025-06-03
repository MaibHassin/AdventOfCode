package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	part := flag.Int("part", 1, "Which puzzle part to solve (1 or 2).")
	flag.Parse()

	input, err := readInput(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *part == 1 {
		fmt.Printf("The number of nice strings: %d\n", NaughtyOrNice(input))
	} else {
		fmt.Printf("The number of nice strings: %d\n", betterNaughtyOrNice(input))
	}

}

func NaughtyOrNice(list []string) int {
	niceStringCount := 0
	for _, line := range list {
		if !containsForbiddenSubstrings(line) {
			if hasMinThreeVowels(line) && hasDoubleLetters(line) {
				niceStringCount += 1
			}
		}
	}
	return niceStringCount
}

func containsForbiddenSubstrings(line string) bool {
	forbiddenSubStrings := []string{"ab", "cd", "pq", "xy"}

	for _, fss := range forbiddenSubStrings {
		if strings.Contains(line, fss) {
			return true
		}
	}
	return false
}

func hasDoubleLetters(line string) bool {
	possibleSubStrings := possibleSubStrings(line)
	for _, ss := range possibleSubStrings {
		if ss[0] == ss[1] {
			return true
		}
	}
	return false
}

func hasMinThreeVowels(line string) bool {
	totalVowels := 0
	vowels := []string{"a", "e", "i", "o", "u"}
	for i := range vowels {
		totalVowels += strings.Count(line, vowels[i])
	}
	return totalVowels >= 3
}

// Part 2
func betterNaughtyOrNice(list []string) int {
	niceStringCount := 0
	for _, line := range list {
		if hasNonOverlappingPairRepeated(line) && hasRepeatingLetterWithOneBetween(line) {
			niceStringCount += 1
		}
	}
	return niceStringCount
}

func hasNonOverlappingPairRepeated(line string) bool {
	possibleSubStrings := possibleSubStrings(line)

	for _, pair := range possibleSubStrings {
		pairs := strings.Count(line, pair)
		if pairs > 1 {
			return true
		}
	}
	return false
}

func hasRepeatingLetterWithOneBetween(line string) bool {
	for i := range len(line) - 2 {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func possibleSubStrings(line string) []string {
	possibleSubStrings := make([]string, 0)

	for i := range len(line) - 1 {
		possibleSubStrings = append(possibleSubStrings, line[i:i+2])
	}
	return possibleSubStrings
}

func readInput(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': %w", filename, err)
	}

	content := strings.Split(string(data), "\n")
	return content, nil
}
