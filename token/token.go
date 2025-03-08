package token

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
	SQUAREROOT
	EXPONENT
)

type Token struct {
	Type  TokenType
	Value string
}
