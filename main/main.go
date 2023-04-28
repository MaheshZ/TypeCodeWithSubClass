package main

import (
	"errors"
	"example/maheshz/typecode/typecode"
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")
	ans, _ := withInterfaceExecute(1, 10, typecode.Add)
	fmt.Printf("With interface: %d\n", ans)
	ans, _ = withoutInterfaceExecute(1, 10, typecode.Add)
	fmt.Printf("Without interface: %d\n", ans)

}


func withInterfaceExecute(op1 int, op2 int, optype int) (int, error) {
	exec, err := typecode.GetOpExecutor(optype)
	if err != nil {
		return 0, nil
	}
	return exec.OperationExecute(op1, op2)

}

func withoutInterfaceExecute(op1 int, op2 int, optype int) (int, error) {
	switch optype {
	case typecode.Add:
		return op1 + op2, nil
	case typecode.Subtract:
		return op1 - op2, nil
	case typecode.Multipy:
		return op1 * op2, nil
	case typecode.Divide:
		return op1 / op2, nil // bad not checking for divide by zero
	default:

		return 0, errors.New(fmt.Sprintf("Unknown optype %d", optype))
	}
}