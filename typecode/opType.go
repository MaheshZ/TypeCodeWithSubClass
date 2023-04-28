package typecode

import (
	"errors"
	"fmt"
)

// int mapping
const (
	Add int = 0
	Subtract     = 1
	Multipy     = 2
	Divide     = 3
)

type OperationExecutor interface{
	OperationExecute(op1 int, opt2 int) (int, error)
}

func GetOpExecutor(optype int) (OperationExecutor, error) {
	switch optype {
	case Add:
		return AddExecutor{SpeedOverAccuracy: false}, nil
	case Subtract:
		return SubExecutor{}, nil
	case Multipy:
		return MultiplyExecutor{}, nil
	case Divide:
		return DivExecutor{CheckDividebyZero: false}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown optype %d", optype))
	}
}