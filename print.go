package main

import (
	"fmt"
	"io"
)

func print(out io.Writer, tokens []Token) {
	for _, token := range tokens {
		fmt.Fprintf(out, "|%-20s| Len => %3d | Value => %s\n", tokenTypeToString(token.T), token.Size, token.Value)
	}
}

func tokenTypeToString(t TokenType) string {
	switch t {
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case LEFT_PARENTHESIS:
		return "LEFT_PARENTHESIS"
	case RIGHT_PARENTHESIS:
		return "RIGHT_PARENTHESIS"
	case LEFT_CURLY_BRACKET:
		return "LEFT_CURLY_BRACKET"
	case RIGHT_CURLY_BRACKET:
		return "RIGHT_CURLY_BRACKET"
	case LEFT_ANGLE_BRACKET:
		return "LEFT_ANGLE_BRACKET"
	case RIGHT_ANGLE_BRACKET:
		return "RIGHT_ANGLE_BRACKET"
	case COLON:
		return "COLON"
	case SEMICOLON:
		return "SEMICOLON"
	case UNDERSCORE:
		return "UNDERSCORE"
	case AMPERSAND:
		return "AMPERSAND"
	case PIPE:
		return "PIPE"
	case COMMA:
		return "COMMA"
	case PERIOD:
		return "PERIOD"
	case EQUALS:
		return "EQUALS"
	case CROSS:
		return "CROSS"
	case HYPHEN:
		return "HYPHEN"
	case LESS:
		return "LESS"
	case GREATER:
		return "GREATER"
	case ASTERISK:
		return "ASTERISK"
	case PERCENTAGE:
		return "PERCENTAGE"
	case SLASH:
		return "SLASH"
	case BACKSLASH:
		return "BACKSLASH"
	case QUOTE:
		return "QUOTE"
	case TICK:
		return "TICK"
	case BACKTICK:
		return "BACKTICK"
	}

	_assert(false, "unreachable")
	return ""
}
