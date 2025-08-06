package main

type Ast struct {
}

type Function struct {
	Statements []Statement
}

type Statement struct {
}

type Identifier struct {
}

type Expression struct {
}

type Addition struct {
	Left  Expression
	Right Expression
}

type Value struct {
}

func parse(tokens []Token) (Ast, error) {

	return Ast{}, nil
}
