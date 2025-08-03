package main

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	STRING TokenType = iota
	NUMBER

	LEFT_PARENTHESIS
	RIGHT_PARENTHESIS
	LEFT_CURLY_BRACKET
	RIGHT_CURLY_BRACKET
	LEFT_ANGLE_BRACKET
	RIGHT_ANGLE_BRACKET

	COLON
	SEMICOLON
	UNDERSCORE
	AMPERSAND
	PIPE
	COMMA
	PERIOD

	EQUALS
	CROSS
	HYPHEN
	LESS
	GREATER
	ASTERISK
	PERCENTAGE

	SLASH
	BACKSLASH

	QUOTE
	TICK
	BACKTICK
)

type Token struct {
	T     TokenType
	Size  int
	Value string
	Line  int
	Pos   int
}

func lex(s string) ([]Token, error) {
	var (
		idx   = 0
		runes = []rune(s)

		line = 1
		pos  = 1
	)

	readWhile := func(p func(b rune) bool) (string, int) {
		start := idx
		end := idx

		for end < len(runes) && p(runes[end]) {
			end++
		}

		return string(runes[start:end]), end - start
	}

	var tokens []Token
	for {
		if idx >= len(runes) {
			break
		}

		r := runes[idx]
		if isWhitespace(r) {
			if isNewline(r) {
				line += 1
				pos = 1
			} else {
				// Jikes...
				if r == '\t' {
					pos += 4
				} else {
					pos += 1
				}
			}

			idx++
			continue
		}

		var token Token
		switch {
		case isLetter(r):
			s, len := readWhile(isAlphaNumOrUnderscore)
			token = Token{T: STRING, Size: len, Value: s}
		case isNumber(r):
			num, len := readWhile(isNumber)
			token = Token{T: NUMBER, Size: len, Value: num}
		case r == '{':
			token = Token{T: LEFT_CURLY_BRACKET, Size: 1, Value: string(r)}
		case r == '}':
			token = Token{T: RIGHT_CURLY_BRACKET, Size: 1, Value: string(r)}
		case r == '(':
			token = Token{T: LEFT_PARENTHESIS, Size: 1, Value: string(r)}
		case r == ')':
			token = Token{T: RIGHT_PARENTHESIS, Size: 1, Value: string(r)}
		case r == '[':
			token = Token{T: LEFT_ANGLE_BRACKET, Size: 1, Value: string(r)}
		case r == ']':
			token = Token{T: RIGHT_ANGLE_BRACKET, Size: 1, Value: string(r)}
		case r == ':':
			token = Token{T: COLON, Size: 1, Value: string(r)}
		case r == ';':
			token = Token{T: SEMICOLON, Size: 1, Value: string(r)}
		case r == '=':
			token = Token{T: EQUALS, Size: 1, Value: string(r)}
		case r == '+':
			token = Token{T: CROSS, Size: 1, Value: string(r)}
		case r == '-':
			token = Token{T: HYPHEN, Size: 1, Value: string(r)}
		case r == '_':
			token = Token{T: UNDERSCORE, Size: 1, Value: string(r)}
		case r == '*':
			token = Token{T: ASTERISK, Size: 1, Value: string(r)}
		case r == '"':
			token = Token{T: QUOTE, Size: 1, Value: string(r)}
		case r == '/':
			token = Token{T: SLASH, Size: 1, Value: string(r)}
		case r == '\\':
			token = Token{T: BACKSLASH, Size: 1, Value: string(r)}
		case r == ',':
			token = Token{T: COMMA, Size: 1, Value: string(r)}
		case r == '<':
			token = Token{T: LESS, Size: 1, Value: string(r)}
		case r == '>':
			token = Token{T: GREATER, Size: 1, Value: string(r)}
		case r == '`':
			token = Token{T: BACKTICK, Size: 1, Value: string(r)}
		case r == '\'':
			token = Token{T: TICK, Size: 1, Value: string(r)}
		case r == '&':
			token = Token{T: AMPERSAND, Size: 1, Value: string(r)}
		case r == '.':
			token = Token{T: PERIOD, Size: 1, Value: string(r)}
		case r == '%':
			token = Token{T: PERCENTAGE, Size: 1, Value: string(r)}
		case r == '|':
			token = Token{T: PIPE, Size: 1, Value: string(r)}
		default:
			return nil, fmt.Errorf("unexpected rune: %s", string(r))
		}
		_assert(token.Size > 0, "unexpected TokenType(%d) = %s with zero size", token.T, token.Value)

		token.Line = line
		token.Pos = pos
		tokens = append(tokens, token)

		idx += token.Size
		pos += token.Size
	}

	return tokens, nil
}

func isAlphaNumOrUnderscore(r rune) bool {
	return isLetter(r) || isNumber(r) || r == '_'
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}

func isNumber(r rune) bool {
	return unicode.IsNumber(r)
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func isNewline(r rune) bool {
	return r == '\n'
}
