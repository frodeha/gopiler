package main

import (
	"fmt"
	"gopiler/lexer"
	"io"
)

func print(out io.Writer, token lexer.Token) {
	fmt.Fprintf(out, "| Token %20s | Length %3d | Line %4d | Pos %3d | Value %s\n", tokenTypeToString(token.T), token.Size, token.Line, token.Pos, token.Value)
}

func tokenTypeToString(t lexer.TokenType) string {
	switch t {
	case lexer.NONE:
		return "<ERROR - NONE>"
	case lexer.STRING:
		return "STRING"
	case lexer.NUMBER:
		return "NUMBER"
	case lexer.LEFT_PARENTHESIS:
		return "LEFT_PARENTHESIS"
	case lexer.RIGHT_PARENTHESIS:
		return "RIGHT_PARENTHESIS"
	case lexer.LEFT_CURLY_BRACKET:
		return "LEFT_CURLY_BRACKET"
	case lexer.RIGHT_CURLY_BRACKET:
		return "RIGHT_CURLY_BRACKET"
	case lexer.LEFT_ANGLE_BRACKET:
		return "LEFT_ANGLE_BRACKET"
	case lexer.RIGHT_ANGLE_BRACKET:
		return "RIGHT_ANGLE_BRACKET"
	case lexer.UNDERSCORE:
		return "UNDERSCORE"
	case lexer.ADD:
		return "CROSS"
	case lexer.SUBTRACT:
		return "HYPHEN"
	case lexer.PACKAGE:
		return "PACKAGE"
	case lexer.IF:
		return "IF"
	case lexer.FOR:
		return "FOR"
	case lexer.FUNC:
		return "FUNC"
	case lexer.IDENTIFIER:
		return "IDENTIFIER"
	case lexer.VAR:
		return "VAR"
	case lexer.DECLARE_INITIALIZE:
		return "DECLARE_INITIALIZE"
	case lexer.ASSIGN:
		return "ASSIGN"
	case lexer.INT:
		return "INT"
	case lexer.RETURN:
		return "RETURN"
	case lexer.NIL:
		return "NIL"
	case lexer.TYPE:
		return "TYPE"
	case lexer.STRUCT:
		return "STRUCT"
	case lexer.INTERFACE:
		return "INTERFACE"
	case lexer.DIVIDE:
		return "DIVIDE"
	}

	_assert(false, "unreachable")
	return ""
}
