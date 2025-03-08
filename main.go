package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ragg967/Calculator/lexer"
	"github.com/ragg967/Calculator/parser"
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
			fmt.Println("Parenthesis: 2 * (3 + 3) = 12\nExponent: 5 ^ 2 = 25\nMultiplication: 5 * 2 = 10\nDivision: 10 / 2 = 5\nAddition: 5 + 5 = 10\nSubtraction: 10 - 5 = 5\nSquare Root: 25 % 2 = 5")
			continue
		}

		result, err := Calculate(input)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Result: %v\n", result)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input:", err)
	}
}
