package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	Assignment = iota + 3
	Complement
	Operation
)

type Signals struct {
	wires map[string]uint16
}

func NewSignals() *Signals {
	return &Signals{
		wires: make(map[string]uint16),
	}
}

func (s *Signal) signalExists() {

}

func (s *Signals) showSignals() {
	fmt.Println(s.wires)
}

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	flag.Parse()

	input, err := readInput(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = input //remove this line
	testInput := []string{"123 -> x", "456 -> y", "x AND y -> d", "x OR y -> e", "x LSHIFT 2 -> f", "y RSHIFT 2 -> g", "NOT x -> h", "NOT y -> i"}

	mySignal := NewSignals()

	for _, t := range testInput {
		parseOperationString(t, mySignal)
	}
	mySignal.showSignals()
}

//  Helper functions

func parseOperationString(expression string, ms *Signals) {
	expressionSlice := strings.Split(expression, " ")

	switch len(expressionSlice) {
	case Assignment:
		parseAssignmentString(expressionSlice, ms)
	case Complement:
		parseComplement(expressionSlice, ms)
	case Operation:
		parseOperation(expressionSlice, ms)
	}
}

func parseOperation(expression []string, ms *Signals) {
	wireKey := expression[4]
	operation := expression[1]

	if slices.Contains(expression, "AND") || slices.Contains(expression, "OR") {
		bitwiseLogicalOperation(uint16(operand1), uint16(operand2), operation, wireKey, ms)
	} else {
		bitwiseShiftOperation(uint16(operand1), uint16(operand2), operation, wireKey, ms)
	}
}

func bitwiseShiftOperation(operand, shiftBy uint16, direction, wireKey string, ms *Signals) {
	switch direction {
	case "LSHIFT":
		ms.wires[wireKey] = operand << shiftBy
	case "RSHIFT":
		ms.wires[wireKey] = operand >> shiftBy
	}
}

func bitwiseLogicalOperation(operand1, operand2 uint16, operation, wireKey string, ms *Signals) {
	switch operation {
	case "AND":
		ms.wires[wireKey] = operand1 & operand2
	case "OR":
		ms.wires[wireKey] = operand1 | operand2
	}
}

func parseComplement(expression []string, ms *Signals) {
	wireKey := expression[3]
	wireComplementKey := expression[1]

	if complementSignal, ok := ms.wires[wireComplementKey]; ok {
		ms.wires[wireKey] = ^complementSignal
	} else {
		fmt.Println("find out the key")
	}
}

func parseAssignmentString(assignmentString []string, ms *Signals) {
	wireName := assignmentString[2]
	wireSignal, err := strconv.ParseUint(assignmentString[0], 10, 16)
	if err != nil {
		fmt.Printf("Error converting value '%s' to integer", assignmentString[0])
		os.Exit(1)
	}
	ms.wires[wireName] = uint16(wireSignal)
}

// Read input file

func readInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file '%s': %w", filename, err)
	}

	content := strings.TrimSpace(string(data))
	return content, nil
}
