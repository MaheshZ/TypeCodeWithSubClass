package typecode

type SubExecutor struct{}

func (addExec SubExecutor) OperationExecute(op1 int, op2 int) (int, error) {
	return op1 - op2, nil
}