package interpreter

import (
	"strconv"
	"strings"
)

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func Calculate(o string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Split(o, " ")

	for _, operatorStr := range operators {

		if isOperator((operatorStr)) {
			right := stack.Pop()
			left := stack.Pop()

			mathFunc := getOperationFunc(operatorStr)

			res := mathFunc(left, right)
			stack.Push(res)

		} else {
			val, err := strconv.Atoi(operatorStr)

			if err != nil {
				return 0, err
			}

			stack.Push(val)
		}
	}

	return int(stack.Pop()), nil
}

func isOperator(o string) bool {
	if o == SUM || o == SUB || o == DIV || o == MUL {
		return true
	}
	return false
}

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}

	}
	return nil
}

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() int {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}

	return 0
}
