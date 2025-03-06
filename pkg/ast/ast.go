package ast

import (
	"fmt"
	"math"

	"github.com/ragg967/GoCalculator/pkg/token"
)

type Node interface {
	Evaluate() (float64, error)
}

type NumberNode struct {
	Value float64
}

func (n NumberNode) Evaluate() (float64, error) {
	return n.Value, nil
}

type BinaryOpNode struct {
	Left     Node
	Right    Node
	Operator token.TokenType
}

func (n BinaryOpNode) Evaluate() (float64, error) {
	leftVal, err := n.Left.Evaluate()
	if err != nil {
		return 0, err
	}

	rightVal, err := n.Right.Evaluate()
	if err != nil {
		return 0, err
	}

	switch n.Operator {
	case token.PLUS:
		return leftVal + rightVal, nil
	case token.MINUS:
		return leftVal - rightVal, nil
	case token.MULTIPLY:
		return leftVal * rightVal, nil
	case token.DIVIDE:
		if rightVal == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return leftVal / rightVal, nil
	case token.EXPONENT:
		return math.Pow(leftVal, rightVal), nil
	case token.SQUAREROOT:
		return math.Pow(leftVal, 1/rightVal), nil
	default:
		return 0, fmt.Errorf("unknown operator")
	}
}
