package parser

import "github.com/exfly/newlang/ast"

var rootAstTree ast.Node

func NewParser(input string) ast.Node {
	lex := NewLex(input)
	exprParse(lex)
	return rootAstTree
}
