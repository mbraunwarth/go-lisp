package lexer

import "fmt"

type Token struct {
	Line  int
	Col   int
	Typ   tokenType
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("token{%d, %d, %s, '%s'}", t.Line, t.Col, t.Typ, t.Value)
}

func makeToken(line, col int, typ tokenType, value string) Token {
	return Token{line, col, typ, value}
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
