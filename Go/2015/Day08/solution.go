package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MaibHassin/AdventOfCode/Go/readfile"
)

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	part := flag.Int("part", 1, "Which part to solve? 1 or 2?")
	flag.Parse()

	input, err := readfile.ReadFileLines(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *part == 1 {
		part1 := DecodedListLength(input)
		fmt.Println("Part 1 answer is", part1)
	} else {
		part2 := EncodedListLength(input)
		fmt.Println("Part 2 answer is", part2)
	}
}

func EncodedListLength(list []string) int {
	originalLen := 0
	encodedLen := 0
	for _, line := range list {
		encodedLen += 2 // for quotes
		encodedLen += strings.Count(line, "\"")
		encodedLen += strings.Count(line, "\\")

		encodedLen += len(line)
		originalLen += len(line)
	}
	return encodedLen - originalLen
}

func DecodedListLength(list []string) int {
	originalLen := 0
	decodedLen := 0
	for _, line := range list {
		decodedLine, err := strconv.Unquote(line)
		if err != nil {
			break
		}
		originalLen += len(line)
		decodedLen += len(decodedLine)
	}
	return originalLen - decodedLen
}
