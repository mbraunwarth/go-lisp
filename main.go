package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
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

	// keep track of current line and column
	line, col := 1, 1

	bufR := bufio.NewReader(input)
	for {
		r, _, err := bufR.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}

		// lex current rune
		switch {
		case r == '(':
			tok := makeToken(line, col, LeftParen, string(r))
			fmt.Println(tok)
		case r == ')':
			tok := makeToken(line, col, RightParen, string(r))
			fmt.Println(tok)
		case r == '+' || r == '-' || r == '*' || r == '/':
			tok := makeToken(line, col, Operator, string(r))
			fmt.Println(tok)
		case unicode.IsLetter(r):
			// keyword vs id
			id := lexID(r, bufR)
			var tok token
			if isKeyword(id) {
				tok = makeToken(line, col, Keyword, id)
			} else {
				tok = makeToken(line, col, ID, id)
			}
			col = col + len(id) - 1
			fmt.Println(tok)
		case r == ';':
			// really a comment
			next, _, err := bufR.ReadRune()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if next != ';' {
				fmt.Println("Unexpected token " + string(r))
				os.Exit(1)
			}

			val := lexComment(bufR)
			tok := makeToken(line, col, Comment, val)

			fmt.Println(tok)
		case unicode.IsNumber(r):
			// int vs float
			i, frac := lexNum(r, bufR) // i and frac are returned as strings

			var tok token
			if frac == "" { // frac == "" means number is integer
				tok = makeToken(line, col, IntLit, i)
				col = col + len(i) - 1
			} else {
				tok = makeToken(line, col, FloatLit, i+"."+frac)
				col = col + len(i+"."+frac) - 1
			}

			fmt.Println(tok)
		case r == '"':
			val := lexString(bufR)
			tok := makeToken(line, col, StringLit, val)

			// manually evolve column, as strings runes (with surrounding quotes)
			// are not counted in `lexString`
			col = col + len(val) + 1

			fmt.Println(tok)
		case r == '\n':
			line = line + 1
			col = 0
		}
		col = col + 1
	}
}

func lexComment(bufR *bufio.Reader) string {
	var sb strings.Builder

	sb.WriteString(";;")

	for {
		next, _, err := bufR.ReadRune()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if next == '\n' {
			// unread newline character as this is needed to reset column counter
			bufR.UnreadRune()
			break
		}

		sb.WriteRune(next)
	}

	return sb.String()
}

func lexID(r rune, bufR *bufio.Reader) string {
	var sb strings.Builder

	sb.WriteRune(r)

	for {
		next, _, err := bufR.ReadRune()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// still in ID?
		if unicode.IsLetter(next) || unicode.IsNumber(next) || next == '_' {
			sb.WriteRune(next)
		} else {
			// unread last rune as this is not part of the id and could belong to another token
			bufR.UnreadRune()
			break
		}
	}

	return sb.String()
}

// lexNum takes the last read rune and the current buffered reader and returns
// two strings. If the read number is an integer ([0-9]+) it is contained in the
// first return value while the second one is an empty string.
// For a floating point number, which is lexically two integer numbers separated
// by a dot `.`, the integer part is contained in the first and the fraction part
// in the second returned value.
func lexNum(r rune, bufR *bufio.Reader) (string, string) {
	var i, frac strings.Builder

	// if the number is a float (i.e. has a fraction part) keep track of it
	hasFrac := false

	// using `currentBuilder` as state to indicate if still lexing integer
	// or fraction
	currentBuilder := &i

	currentBuilder.WriteRune(r)

	for {
		next, _, err := bufR.ReadRune()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if next == '.' {
			// parsing frac part now -> number is float
			hasFrac = true
			currentBuilder = &frac
			continue
		}

		// reached end of number, push lastly read rune back into the reader
		if !unicode.IsNumber(next) {
			bufR.UnreadRune()
			break
		}
		currentBuilder.WriteRune(next)
	}

	// PoC: as there is no global lexer state and the lex functions are not
	// aware of the current global line and column positions, syntax errors
	// are pretty shitty.

	// syntax error, no fraction following a dot
	if hasFrac && frac.Len() == 0 {
		fmt.Println("unexpected character, need at least one number after a `.`")
	}

	return i.String(), frac.String()
}

func lexString(bufR *bufio.Reader) string {
	var sb strings.Builder

	for {
		next, _, err := bufR.ReadRune()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// string literal ends at the next " symbol
		if next == '"' {
			break
		}

		// PoC: ignore errors here
		sb.WriteRune(next)

	}

	return sb.String()
}

func isKeyword(id string) bool {
	return keywords[id]
}
