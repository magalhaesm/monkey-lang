package token

// TokenType represents the type of a token in the language.
type TokenType string

// Token represents a token with its type and literal value.
type Token struct {
	Type    TokenType
	Literal string
}

// Constants for token types.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// keywords maps identifiers to their corresponding token types if they are keywords.
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent checks if the identifier is a keyword and returns the appropriate token type.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
