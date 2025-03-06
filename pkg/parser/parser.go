package parser

import (
	"fmt"
	"strconv"

	"github.com/ragg967/GoCalculator/pkg/ast"
	"github.com/ragg967/GoCalculator/pkg/token"
)

func Parse(tokens []token.Token) (ast.Node, error) {

	var parseExpression func(int) (ast.Node, int, error)
	var parseFactor func(int) (ast.Node, int, error)
	var parseTerm func(int) (ast.Node, int, error)

	parseFactor = func(start int) (ast.Node, int, error) {
		if start >= len(tokens) {
			return nil, 0, fmt.Errorf("unexpected end of expression")
		}

		if tokens[start].Type == token.NUMBER {
			Value, err := strconv.ParseFloat(tokens[start].Value, 64)
			if err != nil {
				return nil, 0, err
			}
			return ast.NumberNode{Value: Value}, start + 1, nil
		}

		if tokens[start].Type == token.LPAREN {
			node, nextIndex, err := parseExpression(start + 1)
			if err != nil {
				return nil, 0, err
			}
			if nextIndex >= len(tokens) || tokens[nextIndex].Type != token.RPAREN {
				return nil, 0, fmt.Errorf("missing closing parenthesis")
			}
			return node, nextIndex + 1, nil
		}

		return nil, 0, fmt.Errorf("unexpected token: %v", tokens[start])
	}

	parseTerm = func(start int) (ast.Node, int, error) {
		leftNode, nextIndex, err := parseFactor(start)
		if err != nil {
			return nil, 0, err
		}

		for nextIndex < len(tokens) {
			if tokens[nextIndex].Type != token.MULTIPLY && tokens[nextIndex].Type != token.DIVIDE {
				break
			}
			operator := tokens[nextIndex].Type
			nextIndex++

			rightNode, nextIndex, err := parseFactor(start)
			if err != nil {
				return nil, 0, err
			}

			leftNode = ast.BinaryOpNode{
				Left:     leftNode,
				Right:    rightNode,
				Operator: operator,
			}
			nextIndex = nextIndex
		}
		return leftNode, nextIndex, nil
	}

	parseExpression = func(start int) (ast.Node, int, error) {

		leftNode, nextIndex, err := parseTerm(start)
		if err != nil {
			return nil, 0, err
		}

		for nextIndex < len(tokens) {
			if tokens[nextIndex].Type != token.PLUS && tokens[nextIndex].Type != token.MINUS {
				break
			}

			operator := tokens[nextIndex].Type
			nextIndex++

			rightNode, newIndex, err := parseTerm(nextIndex)
			if err != nil {
				return nil, 0, err
			}

			leftNode = ast.BinaryOpNode{
				Left:     leftNode,
				Right:    rightNode,
				Operator: operator,
			}
			nextIndex = newIndex
		}
		return leftNode, nextIndex, nil
	}

	node, _, err := parseExpression(0)
	return node, err
}
