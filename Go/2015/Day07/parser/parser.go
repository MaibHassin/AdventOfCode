package parser

import (
	"strconv"
	"strings"

	"github.com/MaibHassin/AdventOfCode/Go/2015/Day07/common"
)

func ParseLine(line string) *common.CircuitOperations {
	op := &common.CircuitOperations{}
	lineSlices := strings.Fields(line)
	op.Wire = lineSlices[len(lineSlices)-1]
	switch len(lineSlices) {
	case 3:
		op.Signal = lineSlices[0]
		op.OpType = common.OpAssignment
	case 4:
		op.Operand1 = lineSlices[1]
		op.OpType = common.OpNot
	case 5:
		op.Operand1 = lineSlices[0]
		switch lineSlices[1] {
		case "AND":
			op.Operand2 = lineSlices[2]
			op.OpType = common.OpAnd
		case "OR":
			op.Operand2 = lineSlices[2]
			op.OpType = common.OpOr
		case "LSHIFT":
			shift, _ := strconv.ParseUint(lineSlices[2], 10, 16)
			op.ShiftAmount = uint16(shift)
			op.OpType = common.OpLShift
		case "RSHIFT":
			shift, _ := strconv.ParseUint(lineSlices[2], 10, 16)
			op.ShiftAmount = uint16(shift)
			op.OpType = common.OpRShift
		}
	}
	return op
}
