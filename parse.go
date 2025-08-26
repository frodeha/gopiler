package main

import (
	"fmt"
	"gopiler/lexer"
)

func parse(l *lexer.Lexer) {
	for {
		token, err := l.Next()
		if err == lexer.ErrEOF {
			break
		}

		if err != nil {
			panic(err)
		}

		switch token.T {
		case lexer.PACKAGE:
			token := must(l.Next())
			_assert(token.T == lexer.IDENTIFIER, "expected identifier after package")
			_assert(token.Value == "main", "expected package identifier to have value 'main'")
			fmt.Printf("Parsed package definition with name '%s'\n", token.Value)

		case lexer.TYPE:
			ident := must(l.Next())
			_assert(ident.T == lexer.IDENTIFIER, "expected identifier after 'type'")

			_type := must(l.Next())
			switch _type.T {
			case lexer.STRUCT:
				consumeExpected(l, lexer.LEFT_CURLY_BRACKET)
				// TODO(frode): parse struct members
				consumeExpected(l, lexer.RIGHT_CURLY_BRACKET)
				fmt.Printf("Parsed struct definition with name '%s'\n", ident.Value)
			case lexer.INTERFACE:
				consumeExpected(l, lexer.LEFT_CURLY_BRACKET)
				// TODO(frode): parse interface members
				consumeExpected(l, lexer.RIGHT_CURLY_BRACKET)
				fmt.Printf("Parsed interface definition with name '%s'\n", ident.Value)
			case lexer.INT:
				// Nothing to do
				fmt.Printf("Parsed integer definition with name '%s'\n", ident.Value)
			default:
				_fatal(fmt.Errorf("invalid type declaration"))
			}

		case lexer.FUNC:
			ident := must(l.Next())
			_assert(ident.T == lexer.IDENTIFIER, "expected identifier after func")

			consumeExpected(l, lexer.LEFT_PARENTHESIS)

		parse_function_arguments:
			for {
				n := must(l.Next())
				switch n.T {
				case lexer.IDENTIFIER:
					consumeExpected(l, lexer.INT)
					// TODO(frode): support more function argument types
				case lexer.RIGHT_PARENTHESIS:
					break parse_function_arguments
				default:
					_fatal(fmt.Errorf("invalid function signature"))
				}
			}

		parse_return_types:
			for {
				n := must(l.Next())
				switch n.T {
				case lexer.ERROR:
					// TODO(frode): support more and multiple return types
				case lexer.LEFT_CURLY_BRACKET:
					break parse_return_types
				default:
					_fatal(fmt.Errorf("invalid function signature, got unexpected %s", tokenTypeToString(n.T)))
				}
			}

		parse_function_body:
			for {
				n := must(l.Next())
				switch n.T {
				// TODO(frode): parse more function body stuff
				case lexer.RETURN:
					must(l.Next()) // Consume a token of "whatever" for now
				case lexer.RIGHT_CURLY_BRACKET:
					break parse_function_body
				default:
					_fatal(fmt.Errorf("invalid function body, got unexpected %s", tokenTypeToString(n.T)))
				}
			}

			fmt.Printf("Parsed function definition with name '%s'\n", ident.Value)
		}
	}
}

func consumeExpected(l *lexer.Lexer, _type lexer.TokenType) {
	token := must(l.Next())
	_assert(token.T == _type, "invalid token: expected %s, but got %s",
		tokenTypeToString(_type),
		tokenTypeToString(token.T),
	)
}

func must(token lexer.Token, err error) lexer.Token {
	if err != nil {
		_fatal(err)
	}
	return token
}
