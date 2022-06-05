package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mbraunwarth/lisp/lexer"
	"github.com/mbraunwarth/lisp/parser"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("At least one file needed as input")
		fmt.Println("Usage: golisp INPUT_FILE")
		os.Exit(1)
	}

	input, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bufR := bufio.NewReader(input)
	ch := make(chan lexer.Token)
	go lexer.Lex(bufR, ch)
	parser.Parse(ch)
}
