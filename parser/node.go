package parser

import "github.com/mbraunwarth/lisp/lexer"

type node struct {
	typ              nodeType
	line, col, depth int
	childs           []*node
}

func makeNode(tok lexer.Token, depth int, typ nodeType) *node {
	return &node{
		typ:    typ,
		line:   tok.Line,
		col:    tok.col,
		depth:  depth,
		childs: make([]*node, 0),
	}
}
