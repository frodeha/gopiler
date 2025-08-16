package main

import (
	"fmt"
	"gopiler/lexer"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Incorrect usage: missing input file")
		os.Exit(1)
	}

	bytes, err := os.ReadFile(os.Args[1])
	_fatal(err)

	_lexer := lexer.NewFromString(string(bytes))
	for {
		token, err := _lexer.Next()
		if err == lexer.ErrEOF {
			break
		}

		if err != nil {
			fmt.Printf("[Error] %s", err)
			return
		}

		print(os.Stdout, token)
	}
}

func _fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func _assert(b bool, msg string, args ...interface{}) {
	if !b {
		panic(fmt.Sprintf("[Assert failed]: %s", fmt.Sprintf(msg, args...)))
	}
}

func _debug(format string, args ...interface{}) {
	if debug := os.Getenv("DEBUG"); debug != "" {
		fmt.Print("[DEBUG] ")
		fmt.Printf(format, args...)
	}
}
