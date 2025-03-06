package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	fmt.Println("GoCalculator - Enter an expression ('exit' to quit and 'operators' to show all operators and what they do):")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			return
		}

		if strings.ToLower(input) == "operators" {
			fmt.Println("Addition: +\nSubtraction: -\nMultiplication: *\nDivision: /\nExponent: ^\nSquare Root: %")
			return
		}

		result, err := Calculate(input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Result: %v\n", result)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input:", err)
	}
}
