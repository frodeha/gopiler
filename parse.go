package main

import "gopiler/lexer"

func parse(l *lexer.Lexer) {
	for {
		token, err := l.Next()
		if err == lexer.ErrEOF {
			break
		}

		if err != nil {
			panic(err)
		}

		switch token.T{
			case lexer.PACKAGE:
				token, err := l.Next()
				if err != nil {
					panic(err)
				}
				_assert(token.T == lexer.IDENTIFIER, "expected identifier after package token")
				_assert(token.Value == "main", "expected package identifier to have value 'main'")
		}
	}
}
