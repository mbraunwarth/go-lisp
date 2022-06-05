package parser

import (
	"fmt"

	"github.com/mbraunwarth/lisp/lexer"
)

func Parse(ch <-chan lexer.Token) {
	for tok := range ch {
		fmt.Println(tok)
	}
}
