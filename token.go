package main

import "fmt"

type token struct {
	line  int
	col   int
	typ   tokenType
	value string
}

func (t token) String() string {
	return fmt.Sprintf("token{%d, %d, %s, '%s'", t.line, t.col, t.typ, t.value)
}

func makeToken(line, col int, typ tokenType, value string) token {
	return token{line, col, typ, value}
}

type tokenType int

const (
	Undefined tokenType = iota
	LeftParen
	RightParen
	Operator
	Keyword
	ID
	Comment

	IntLit
	FloatLit
	StringLit
)

func (t tokenType) String() string {
	switch t {
	case Undefined:
		return "Undefined"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case Operator:
		return "Operator"
	case Keyword:
		return "Keyword"
	case ID:
		return "ID"
	case Comment:
		return "Comment"
	case IntLit:
		return "IntLit"
	case FloatLit:
		return "FloatLit"
	case StringLit:
		return "StringLit"
	default:
		return "ERROR"
	}
}
