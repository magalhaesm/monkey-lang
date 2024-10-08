package lexer

import "github.com/magalhaesm/monkey-lang/token"

// Lexer represents the lexer (or scanner) for tokenizing the input string.
type Lexer struct {
	input        string // The input string (source code) to tokenize
	position     int    // Current position in input (points to current character)
	readPosition int    // Current reading position in input (points to next character)
	ch           byte   // Current character under examination
}

// New initializes a new Lexer for the given input string.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar advances the lexer to the next character in the input.
// Sets l.ch to 0 if the end of the input is reached (EOF).
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken retrieves the next token from the input and advances the lexer.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newTwoCharToken(token.EQ, ch, l.ch)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newTwoCharToken(token.NOT_EQ, ch, l.ch)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readNumber reads an integer from the input and advances the lexer's position.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isDigit checks if the given character is a digit ('0' to '9').
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// skipWhitespace skips over whitespace characters like spaces, tabs, and newlines.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// readIdentifier reads an identifier (a sequence of letters or underscores) from the input.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter checks if the given character is a letter (either 'a'-'z', 'A'-'Z') or an underscore ('_').
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// newToken creates a new token with the given type and literal value.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// newTwoCharToken creates a new token with the given type and a literal value
// formed by concatenating two characters (ch1 and ch2).
func newTwoCharToken(tokenType token.TokenType, ch1 byte, ch2 byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch1) + string(ch2),
	}
}
