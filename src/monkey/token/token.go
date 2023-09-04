package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// NO square bracket `[]` ?

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	EQ     = "=="
	NOT_EQ = "!="

	BANG     = "!"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	LT = "<"
	GT = ">"
)

type TokenType string

// Token represents a token type and its literal value.
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
