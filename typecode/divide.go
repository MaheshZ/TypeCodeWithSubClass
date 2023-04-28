package typecode

type DivExecutor struct {
	CheckDividebyZero bool
}

func (divExec DivExecutor) OperationExecute(op1 int, op2 int) (int, error) {
	return op1 / op2, nil
}