package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MaibHassin/AdventOfCode/Go/2015/Day07/parseinstructions"
	"github.com/MaibHassin/AdventOfCode/Go/readfile"
)

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	flag.Parse()

	input, err := readfile.ReadFileLines(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, inst := range input {
		parsedInstruction, err := parseinstructions.ParseLine(inst)
		_ = err
		fmt.Printf("parsedInstruction.Wire: %v\n", parsedInstruction.Wire)
		fmt.Printf("parsedInstruction.Signal: %v\n", parsedInstruction.Signal)
		fmt.Printf("parsedInstruction.Operation: %v\n", parsedInstruction.Operation)
		fmt.Printf("parsedInstruction.Operand1: %v\n", parsedInstruction.Operand1)
		fmt.Printf("parsedInstruction.Operand2: %v\n", parsedInstruction.Operand2)
		fmt.Printf("parsedInstruction.Shift: %v\n", parsedInstruction.Shift)
		fmt.Println("================================================================")
	}
}
