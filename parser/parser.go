package parser

import (
	"fmt"
	"strconv"

	"github.com/ragg967/Calculator/ast"
	"github.com/ragg967/Calculator/token"
)

func Parse(tokens []token.Token) (ast.Node, error) {
	// Define parsing functions using mutual recursion
	var parseExpression, parseTerm, parseFactor, parsePrimary func(int) (ast.Node, int, error)

	// Primary: numbers and parenthesized expressions
	parsePrimary = func(start int) (ast.Node, int, error) {
		if start >= len(tokens) {
			return nil, 0, fmt.Errorf("unexpected end of expression")
		}

		if tokens[start].Type == token.NUMBER {
			value, err := strconv.ParseFloat(tokens[start].Value, 64)
			if err != nil {
				return nil, 0, err
			}
			return ast.NumberNode{Value: value}, start + 1, nil
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

	// Factor: exponents and square roots
	parseFactor = func(start int) (ast.Node, int, error) {
		leftNode, nextIndex, err := parsePrimary(start)
		if err != nil {
			return nil, 0, err
		}

		for nextIndex < len(tokens) {
			if tokens[nextIndex].Type != token.EXPONENT && tokens[nextIndex].Type != token.SQUAREROOT {
				break
			}

			operator := tokens[nextIndex].Type
			nextIndex++

			rightNode, newNextIndex, err := parsePrimary(nextIndex)
			if err != nil {
				return nil, 0, err
			}

			leftNode = ast.BinaryOpNode{
				Left:     leftNode,
				Right:    rightNode,
				Operator: operator,
			}
			nextIndex = newNextIndex
		}
		return leftNode, nextIndex, nil
	}

	// Term: multiplication and division
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

			rightNode, newNextIndex, err := parseFactor(nextIndex)
			if err != nil {
				return nil, 0, err
			}

			leftNode = ast.BinaryOpNode{
				Left:     leftNode,
				Right:    rightNode,
				Operator: operator,
			}
			nextIndex = newNextIndex
		}
		return leftNode, nextIndex, nil
	}

	// Expression: addition and subtraction
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

			rightNode, newNextIndex, err := parseTerm(nextIndex)
			if err != nil {
				return nil, 0, err
			}

			leftNode = ast.BinaryOpNode{
				Left:     leftNode,
				Right:    rightNode,
				Operator: operator,
			}
			nextIndex = newNextIndex
		}

		return leftNode, nextIndex, nil
	}

	node, nextIndex, err := parseExpression(0)
	if err != nil {
		return nil, err
	}

	if nextIndex < len(tokens) {
		return nil, fmt.Errorf("unexpected token at end: %v", tokens[nextIndex])
	}

	return node, nil
}
