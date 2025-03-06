package main

import (
	"fmt"

	"github.com/ragg967/GoCalculator/pkg/lexer"
	"github.com/ragg967/GoCalculator/pkg/parser"
)

func Calculate(expression string) (float64, error) {

	tokens, err := lexer.Tokenize(expression)
	if err != nil {
		return 0, err
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		return 0, err
	}

	result, err := ast.Evaluate()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	result, err := Calculate("3 + 3 * 2")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
