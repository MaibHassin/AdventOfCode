package common

type OperationType int

const (
	OpAssignment OperationType = iota
	OpNot
	OpAnd
	OpOr
	OpLShift
	OpRShift
)

type CircuitOperations struct {
	Wire        string
	OpType      OperationType
	Operand1    string
	Operand2    string
	Signal      string
	ShiftAmount uint16
}

type SignalStore struct {
	Signals map[string]uint16
}
