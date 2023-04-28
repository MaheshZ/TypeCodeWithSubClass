package typecode

type MultiplyExecutor struct {
	SpeedOverAccuracy bool
}

func (mulExec MultiplyExecutor) OperationExecute(op1 int, op2 int) (int, error) {
	return op1 * op2, nil
}