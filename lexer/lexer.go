package lexer

import (
	"errors"
	"fmt"
	"unicode"
)

var (
	ErrEOF = errors.New("no more tokens")
)

type Lexer struct {
	idx   int
	runes []rune

	line     int
	position int
}

func (l *Lexer) All() ([]Token, error) {
	var tokens []Token
	for {
		token, err := l.Next()
		if err == ErrEOF {
			break
		}

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (l *Lexer) Next() (Token, error) {
	l.consumeWhitespace()

	r, ok := l.peek()
	if !ok {
		return Token{}, ErrEOF
	}

	token := Token{
		T:    NONE,
		Line: l.line,
		Pos:  l.position,
	}

	// Keywords and Identifiers
	if unicode.IsLetter(r) {
		s, size := l.captureWhile(func(r rune) bool { return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_' })
		if tokenType, ok := keywords[s]; ok {
			token.T = tokenType
		} else {
			token.T = IDENTIFIER
		}
		token.Value = s
		token.Size = size
		return token, nil
	}

	// Numbers
	if unicode.IsNumber(r) {
		token.T = NUMBER
		token.Value, token.Size = l.captureWhile(func(r rune) bool { return unicode.IsNumber(r) || r == '.' })
		return token, nil
	}

	// String literals
	if r == '"' {
		l.consume() // Skip "
		token.T = STRING
		token.Value, token.Size = l.captureWhile(func(r rune) bool { return r != '"' })

		r, ok := l.peek() // Skip "
		if !ok {
			return Token{}, fmt.Errorf("parse error: failed to parse string: expected '\"' at end of string, but got EOF")
		}
		if r != '"' {
			return Token{}, fmt.Errorf("parse error: failed to parse string: expected '\"' at end of string, but got %q", r)
		}
		l.consume()
		return token, nil
	}

	// Declaration initialization
	rs, ok := l.peekN(2)
	if ok && rs[0] == ':' && rs[1] == '=' {
		token.T = DECLARE_INITIALIZE
		token.Value = string(rs)
		token.Size = 2
		l.consumeN(2)
		return token, nil
	}

	switch r {
	case '{':
		token.T = LEFT_CURLY_BRACKET
	case '}':
		token.T = RIGHT_CURLY_BRACKET
	case '(':
		token.T = LEFT_PARENTHESIS
	case ')':
		token.T = RIGHT_PARENTHESIS
	case '[':
		token.T = LEFT_ANGLE_BRACKET
	case ']':
		token.T = RIGHT_ANGLE_BRACKET
	case '=':
		token.T = ASSIGN
	case '_':
		token.T = UNDERSCORE
	case '+':
		token.T = ADD
	case '-':
		token.T = SUBTRACT
	case '*':
		token.T = MULTIPLY_ASTERISK
	case '/':
		token.T = DIVIDE
	default:
		return Token{}, fmt.Errorf("unexpected rune %q on line %d position %d", r, l.line, l.position)
	}

	token.Size = 1
	token.Value = string(r)
	l.consume()

	return token, nil
}

var keywords = map[string]TokenType{
	"package":   PACKAGE,
	"func":      FUNC,
	"if":        IF,
	"for":       FOR,
	"var":       VAR,
	"int":       INT,
	"return":    RETURN,
	"nil":       NIL,
	"type":      TYPE,
	"struct":    STRUCT,
	"interface": INTERFACE,
}

func (l *Lexer) peek() (rune, bool) {
	if l.idx >= len(l.runes) {
		return rune(0), false
	}

	return l.runes[l.idx], true
}

func (l *Lexer) peekN(n int) ([]rune, bool) {
	if l.idx+n >= len(l.runes) {
		// Protect from out of bounds
		return l.runes[l.idx:], false
	}
	return l.runes[l.idx : l.idx+n], true
}

func (l *Lexer) captureWhile(p func(b rune) bool) (string, int) {
	var (
		start = l.idx
		end   = l.idx
	)
	for {
		r, ok := l.peek()
		if !ok || !p(r) {
			break
		}
		l.consume()
		end++
	}
	capture := string(l.runes[start:end])
	size := end - start
	return capture, size
}

func (l *Lexer) consumeN(n int) {
	for idx := 0; idx < n; idx++ {
		l.consume()
	}
}

func (l *Lexer) consume() {
	r, ok := l.peek()
	if !ok {
		return
	}

	// Maintain file positions
	if r == '\n' || r == '\r' {
		l.line += 1
		l.position = 1
	} else {
		// Jikes...
		if r == '\t' {
			l.position += 4
		} else {
			l.position += 1
		}
	}

	l.idx += 1
}

func (l *Lexer) consumeWhitespace() {
	for {
		r, ok := l.peek()
		if !ok {
			return
		}
		if unicode.IsSpace(r) {
			l.consume()
			continue
		}

		// Comments
		// TODO(frode): Support multi-line comments
		rs, ok := l.peekN(2)
		if ok && rs[0] == '/' && rs[1] == '/' {
			l.captureWhile(func(b rune) bool { return b != '\n' })
			continue
		}

		break
	}
}

func NewFromString(s string) *Lexer {
	return &Lexer{
		idx:      0,
		runes:    []rune(s),
		line:     1,
		position: 1,
	}
}
