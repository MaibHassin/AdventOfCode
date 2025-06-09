package resolver

import (
	"strconv"

	"github.com/MaibHassin/AdventOfCode/Go/2015/Day07/common"
)

func NewSignalStore() *common.SignalStore {
	return &common.SignalStore{
		Signals: make(map[string]uint16),
	}
}

func HasSignal(wire string, ss *common.SignalStore) (uint16, bool) {
	signal, ok := ss.Signals[wire]
	return signal, ok
}

func ResolveOperation(op common.CircuitOperations, ss *common.SignalStore) {
	switch op.OpType {
	case common.OpAssignment:
		signal, ok := isNum(op.Signal, ss)
		if ok {
			assignSignal(op.Wire, signal, ss)
		}
	case common.OpNot:
		calculateAndAssignComplement(op.Wire, op.Operand1, ss)
	case common.OpAnd, common.OpOr:
		calculateAndAssignLogic(op.Wire, op.Operand1, op.Operand2, op.OpType, ss)
	case common.OpLShift, common.OpRShift:
		calculateAndAssignShift(op.Wire, op.Operand1, op.ShiftAmount, op.OpType, ss)
	}
}

func isNum(val string, ss *common.SignalStore) (uint16, bool) {
	valInt, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		valInt, ok := HasSignal(val, ss)
		if ok {
			return valInt, ok
		}
		return valInt, ok
	}
	return uint16(valInt), true
}

func assignSignal(wire string, signal uint16, ss *common.SignalStore) {
	ss.Signals[wire] = signal
}

func calculateAndAssignComplement(wire, operand string, ss *common.SignalStore) {
	if _, ok := HasSignal(operand, ss); ok {
		signal := ^ss.Signals[operand]
		assignSignal(wire, signal, ss)
	}
}

func calculateAndAssignLogic(wire, op1, op2 string, logic common.OperationType, ss *common.SignalStore) {
	op1Val, op1Exists := isNum(op1, ss)
	op2Val, op2Exists := isNum(op2, ss)

	if op1Exists && op2Exists {
		var signal uint16
		switch logic {
		case common.OpAnd:
			signal = op1Val & op2Val
		case common.OpOr:
			signal = op1Val | op2Val
		}
		assignSignal(wire, signal, ss)
	}
}

func calculateAndAssignShift(wire, operand string, shiftamount uint16, shiftdirection common.OperationType, ss *common.SignalStore) {
	if operandSignal, ok := HasSignal(operand, ss); ok {
		var signal uint16
		switch shiftdirection {
		case common.OpLShift:
			signal = operandSignal << shiftamount
		case common.OpRShift:
			signal = operandSignal >> shiftamount
		}
		assignSignal(wire, signal, ss)
	}
}
