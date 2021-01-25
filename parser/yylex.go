//go:generate goyacc -o expr.go -p "expr" expr.y
package parser

import (
	"log"

	"github.com/exfly/newlang/lexer"
	"github.com/exfly/newlang/token"
)

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const EOF = 0

func NewLex(input string) *exprLex {
	return &exprLex{
		Lexer: lexer.New(input),
	}
}

type exprLex struct {
	*lexer.Lexer
}

func (x *exprLex) Lex(yylval *exprSymType) int {
	next := x.NextToken()

	log.Printf("%+v\n", next)

	switch next.Type {
	case token.EOF:
		return EOF
	// case token.EOL:
	// 	return EOL
	case token.COMMA:
		yylval.tok = next
		return ','
	case token.PLUS:
		yylval.tok = next
		return '+'
	case token.MINUS:
		yylval.tok = next
		return '-'
	case token.MULTIPLY:
		yylval.tok = next
		return '*'
	case token.DIVIDE:
		yylval.tok = next
		return '/'
	case token.LPAREN:
		yylval.tok = next
		return '('
	case token.RPAREN:
		yylval.tok = next
		return ')'
	case token.LBRACE:
		yylval.tok = next
		return '{'
	case token.RBRACE:
		yylval.tok = next
		return '}'
	case token.SEMICOLON:
		yylval.tok = next
		return ';'
	case token.GT:
		yylval.tok = next
		return '>'
	case token.LT:
		yylval.tok = next
		return '<'
	case token.EQ:
		yylval.tok = next
		return EQ
	case token.ASSIGN:
		yylval.tok = next
		return ASSIGN

	case token.BIND:
		yylval.tok = next
		return BIND

	case token.IF:
		yylval.tok = next
		return IF
	case token.ELSE:
		yylval.tok = next
		return ELSE

	case token.FUNCTION:
		yylval.tok = next
		return FUNCTION
	case token.RETURN:
		yylval.tok = next
		return RETURN

	case token.IDENT:
		yylval.tok = next
		return IDENTIFIER

	case token.INT:
		yylval.tok = next
		return NUMBER

	case token.ILLEGAL:
		yylval.tok = next
		return ILLIGAL
	default:
		log.Printf("unrecognized character %+v", next)
	}

	return -1
}

// The parser calls this method on a parse error.
func (x *exprLex) Error(s string) {
	log.Printf("parse error: %s", s)
}
