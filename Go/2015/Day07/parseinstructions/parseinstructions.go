package parseinstructions

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	assignment = iota + 3
	complement
	connection
)

// Instruction represents a circuit instruction.
//
// For assignment instructions (e.g., "123 -> x"), it stores the target wire (Wire) and the signal value (Signal).
// Operation and operands are empty.
//
// For complement instructions (e.g., "NOT a -> x"), it stores the target wire (Wire), the operation "NOT" (Operation),
// and the operand to be complemented (Operand1). Signal and Operand2 are empty.
//
// For connection instructions (e.g., "a AND b -> x", "a OR b -> x", "a LSHIFT 2 -> x", "a RSHIFT 2 -> x"),
// it stores the target wire (Wire), the operation ("AND", "OR", "LSHIFT", or "RSHIFT") (Operation),
// and the two operands (Operand1 and Operand2). Signal is empty.
type Instruction struct {
	Wire      string
	Signal    uint16
	Operation string
	Operand1  string
	Operand2  string
	Shift     uint16
}

func ParseLine(line string) (Instruction, error) {
	lineSlince := strings.Split(line, " ")

	switch len(lineSlince) {
	case assignment:
		return parseAssignment(lineSlince)
	case complement:
		return parseComplement(lineSlince)
	case connection:
		return parseConnection(lineSlince)
	default:
		return Instruction{}, fmt.Errorf("unrecognized instruction format: %s", line)
	}
}

func parseAssignment(line []string) (Instruction, error) {
	signal, err := strconv.ParseUint(line[0], 10, 16)

	if err != nil {
		return Instruction{}, fmt.Errorf("invalid signal value '%s': %w", line[0], err)
	}

	return Instruction{
		Wire:      line[2],
		Signal:    uint16(signal),
		Operation: "SIGNAL ASSIGNMENT",
	}, nil
}

func parseComplement(line []string) (Instruction, error) {
	return Instruction{
		Wire:      line[3],
		Operation: line[0],
		Operand1:  line[1],
	}, nil
}

func parseConnection(line []string) (Instruction, error) {
	inst := Instruction{
		Wire:      line[4],
		Operation: line[1],
		Operand1:  line[0],
	}

	switch inst.Operation {
	case "AND", "OR":
		inst.Operand2 = line[2]
	case "LSHIFT", "RSHIFT":
		shiftVal, err := strconv.ParseUint(line[2], 10, 16)
		if err != nil {
			return Instruction{}, fmt.Errorf("invalid shift value '%s' for %s operation: %w", line[2], inst.Operation, err)
		}
		inst.Shift = uint16(shiftVal)
	default:
		return Instruction{}, fmt.Errorf("unknown connection operation: %s", inst.Operation)
	}
	return inst, nil
}
