package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MaibHassin/AdventOfCode/Go/2015/Day07/common"
	"github.com/MaibHassin/AdventOfCode/Go/2015/Day07/parser"
	"github.com/MaibHassin/AdventOfCode/Go/2015/Day07/resolver"
	"github.com/MaibHassin/AdventOfCode/Go/readfile"
)

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	requiredWire := flag.String("wire", "a", "Signal on which wire?")
	part := flag.Int("part", 1, "Which part to run? 1 or 2?")
	flag.Parse()

	input, err := readfile.ReadFileLines(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	signalStore := resolver.NewSignalStore()
	circuitOperations := []common.CircuitOperations{}
	for _, line := range input {
		parsedInstruction := *parser.ParseLine(line)
		circuitOperations = append(circuitOperations, parsedInstruction)
	}

	part1 := getSignal(circuitOperations, signalStore, *requiredWire)

	if *part == 2 {
		var part2 uint16
		signalStorePart2 := resolver.NewSignalStore()
		signalStorePart2.Signals["b"] = part1
		part2 = getSignal(circuitOperations, signalStorePart2, *requiredWire)
		fmt.Println(part2)
	} else {
		fmt.Println(part1)
		fmt.Println(len(signalStore.Signals))
	}

}

func getSignal(co []common.CircuitOperations, ss *common.SignalStore, whichWire string) uint16 {
	for {
		for _, op := range co {
			_, ok := resolver.HasSignal(op.Wire, ss)
			if !ok {
				resolver.ResolveOperation(op, ss)
			}
		}

		if _, ok := ss.Signals[whichWire]; ok {
			break
		}
	}
	return ss.Signals[whichWire]
}
