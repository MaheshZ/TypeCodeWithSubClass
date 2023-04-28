package typecode

type AddExecutor struct {
	SpeedOverAccuracy bool
}

func (addExec AddExecutor) OperationExecute(op1 int, op2 int) (int, error) {
	return op1 + op2, nil
}