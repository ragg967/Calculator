package lexer

import (
	"fmt"
	"unicode"

	"github.com/ragg967/GoCalculator/pkg/token"
)

func Tokenize(expression string) ([]token.Token, error) {
	tokens := []token.Token{}
	currentNumber := ""

	for _, char := range expression {
		switch {
		case unicode.IsDigit(char) || char == '.':
			currentNumber += string(char)
		case char == '+':
			if currentNumber != "" {
				tokens = append(tokens, token.Token{Type: token.NUMBER, Value: currentNumber})
				currentNumber = ""
			}
			tokens = append(tokens, token.Token{Type: token.PLUS, Value: "+"})
		case char == '-':
			if currentNumber != "" {
				tokens = append(tokens, token.Token{Type: token.NUMBER, Value: currentNumber})
				currentNumber = ""
			}
			tokens = append(tokens, token.Token{Type: token.MINUS, Value: "-"})
		case char == '*':
			if currentNumber != "" {
				tokens = append(tokens, token.Token{Type: token.NUMBER, Value: currentNumber})
				currentNumber = ""
			}
			tokens = append(tokens, token.Token{Type: token.MULTIPLY, Value: "*"})
		case char == '/':
			if currentNumber != "" {
				tokens = append(tokens, token.Token{Type: token.NUMBER, Value: currentNumber})
				currentNumber = ""
			}
			tokens = append(tokens, token.Token{Type: token.DIVIDE, Value: "/"})
		case unicode.IsSpace(char):
			if currentNumber != "" {
				tokens = append(tokens, token.Token{Type: token.NUMBER, Value: currentNumber})
				currentNumber = ""
			}
		default:
			return nil, fmt.Errorf("unexpected character: %c", char)
		}
	}

	if currentNumber != "" {
		tokens = append(tokens, token.Token{Type: token.NUMBER, Value: currentNumber})
	}

	return tokens, nil
}
