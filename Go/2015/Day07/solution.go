package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MaibHassin/AdventOfCode/Go/readfile"
)

type CircuitOperations struct {
	Wire        string
	Signal      string
	Operation   string
	Operand1    string
	Operand2    string
	ShiftAmount string
}

func NewCircuitOperations() *CircuitOperations {
	return &CircuitOperations{}
}

type SignalStore struct {
	Signals map[string]uint16
}

func NewSignalStore() *SignalStore {
	return &SignalStore{
		Signals: make(map[string]uint16),
	}
}

func (ss *SignalStore) addSignal(key, signal string) {
	signalToInt16, err := strconv.ParseUint(signal, 10, 16)
	if err != nil {
		return
	}

	ss.Signals[key] = uint16(signalToInt16)
}

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	part := flag.Int("part", 1, "Which part to run? 1 or 2?")
	flag.Parse()

	input, err := readfile.ReadFileLines(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	circuitOperations := []CircuitOperations{}
	for _, line := range input {
		co := NewCircuitOperations()
		co.classifyOperations(line)
		circuitOperations = append(circuitOperations, *co)
	}

	finalResult := NewSignalStore()

	for _, co := range circuitOperations {
		if co.Operation == "ASSIGNMENT" {
			initialAssignment(co, finalResult)
		}
		if *part == 2 {
			finalResult.Signals["b"] = uint16(16076)
		}
	}

	loopRan := 0
	for {
		if len(finalResult.Signals) == len(input) {
			break
		}
		loopRan += 1

		for _, op := range circuitOperations {
			switch op.Operation {
			case "ASSIGNMENT":
				if finalResult.hasSignal(op.Wire) {
					continue
				}
				signalVal, err := strconv.ParseUint(op.Signal, 10, 16)
				if err == nil {
					finalResult.Signals[op.Wire] = uint16(signalVal)
				} else {
					if finalResult.hasSignal(op.Signal) {
						finalResult.Signals[op.Wire] = finalResult.Signals[op.Signal]
					}
				}

			case "NOT":
				if finalResult.hasSignal(op.Operand1) {
					complement(op.Wire, op.Operand1, finalResult)
				}
			case "LSHIFT", "RSHIFT":
				if finalResult.hasSignal(op.Operand1) {
					shift(op.Wire, op.Operand1, op.ShiftAmount, op.Operation, finalResult)
				}
			case "AND", "OR":
				logic(op.Wire, op.Operand1, op.Operand2, op.Operation, finalResult)
			}
		}

	}
	fmt.Printf("finalResult: %v\n", finalResult.Signals["a"])
}

func logic(wire, operandA, operandB, logic string, fr *SignalStore) {
	// if both are numbers
	if isNum(operandA) && isNum(operandB) {
		numA, _ := strconv.ParseUint(operandA, 10, 16)
		numB, _ := strconv.ParseUint(operandB, 10, 16)
		assignLogicOperations(wire, logic, uint16(numA), uint16(numB), fr)
		return
	}
	// if a is num and b exists
	if isNum(operandA) && fr.hasSignal(operandB) {
		numA, _ := strconv.ParseUint(operandA, 10, 16)
		numB := fr.Signals[operandB]
		assignLogicOperations(wire, logic, uint16(numA), numB, fr)
		return
	}
	// if a exists and b is num
	if isNum(operandB) && fr.hasSignal(operandA) {
		numA := fr.Signals[operandA]
		numB, _ := strconv.ParseUint(operandB, 10, 16)
		assignLogicOperations(wire, logic, numA, uint16(numB), fr)
		return
	}
	// if both exists
	if fr.hasSignal(operandA) && fr.hasSignal(operandB) {
		numA := fr.Signals[operandA]
		numB := fr.Signals[operandB]
		assignLogicOperations(wire, logic, numA, numB, fr)
		return
	}
}

func assignLogicOperations(wire, operation string, opa, opb uint16, fr *SignalStore) {
	switch operation {
	case "AND":
		fr.Signals[wire] = opa & opb
	case "OR":
		fr.Signals[wire] = opa | opb
	}
}

func isNum(op string) bool {
	_, err := strconv.ParseUint(op, 10, 16)
	return err == nil
}

func shift(wire, operand, shiftVal, shiftDirection string, fr *SignalStore) {
	shiftValToInt16, err := strconv.ParseUint(shiftVal, 10, 16)
	if err != nil {
		return
	}
	if shiftDirection == "LSHIFT" {
		fr.Signals[wire] = fr.Signals[operand] << uint16(shiftValToInt16)
	} else {
		fr.Signals[wire] = fr.Signals[operand] >> uint16(shiftValToInt16)
	}
}

func complement(wire, operand string, fr *SignalStore) {
	fr.Signals[wire] = ^fr.Signals[operand]
}

func (ss *SignalStore) hasSignal(wire string) bool {
	_, ok := ss.Signals[wire]

	return ok
}

func initialAssignment(co CircuitOperations, fr *SignalStore) {
	signalToInt16, err := strconv.ParseUint(co.Signal, 10, 16)
	if err != nil {
		return
	}
	fr.Signals[co.Wire] = uint16(signalToInt16)
}

func (co *CircuitOperations) classifyOperations(ops string) {

	opS := strings.Split(ops, " ")
	switch len(opS) {
	case 3:
		co.Wire = opS[2]
		co.Signal = opS[0]
		co.Operation = "ASSIGNMENT"
	case 4:
		co.Wire = opS[3]
		co.Operand1 = opS[1]
		co.Operation = opS[0]
	case 5:
		if strings.Contains(ops, "LSHIFT") || strings.Contains(ops, "RSHIFT") {
			co.Wire = opS[4]
			co.Operation = opS[1]
			co.Operand1 = opS[0]
			co.ShiftAmount = opS[2]
		} else {
			co.Wire = opS[4]
			co.Operation = opS[1]
			co.Operand1 = opS[0]
			co.Operand2 = opS[2]
		}
	}
}
