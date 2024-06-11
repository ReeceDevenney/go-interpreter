package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Indentifiers + literals
	IDENT = "IDENT" //add, foobar, x, yk
	INT   = "INT"
	//operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
