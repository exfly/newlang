// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an example of a goyacc program.
// To build it:
// goyacc -p "expr" expr.y (produces y.go)
// go build -o expr y.go
// expr
// > <type an expression>

%{

package parser

import (
	"github.com/exfly/newlang/token"
	"github.com/exfly/newlang/ast"
)
%}

%union {
	astTree ast.Node
	tok token.Token
	statement ast.Statement
	expr ast.Expression
	block *ast.BlockStatement
	id []*ast.Identifier
	exprs []ast.Expression
}

%token '+' '-' '*' '/' '(' ')' ';' ','
%token ILLIGAL
%token <tok> NUMBER IDENTIFIER BIND ASSIGN FUNCTION EQ
%token <tok> IF ELSE WHILE RETURN
%token EOL
%type <astTree>  program
%type <expr> bindexpr assignexpr funcexpr expr ident ifexpr factor term number
%type <statement> stmt retstmt
%type <tok> '+' '-' '*' '/' '(' ')' ',' '{' '}' IF ELSE
%type <tok> comparison '>' '<' '='
%type <tok> end_of_stmt ';' EOL
%type <block> stmts block

%type <id> funcargs
%type <exprs> callargs

%start program

%%

program: stmts { rootAstTree = $1 }
	;

stmts: stmt end_of_stmt { $$ = ast.NewRowBlockStatement(token.Token{Type: token.LBRACE, Literal: "{"}, nil); $$.AppendStatement($1) }
	| stmts stmt end_of_stmt { $1.AppendStatement($2)}
	;

end_of_stmt: ';' | EOL

stmt: expr { $$ = ast.NewExpressionStatement(token.Token{Literal: "?"}, $1.(ast.Expression)) }
	| retstmt { $$ = $1 }
	;

block: '{' stmts '}' { $$ = $2 }
	| '{' '}' { $$ = ast.NewRowBlockStatement(token.Token{Type: token.LBRACE, Literal: "{"}, nil) }
	;

bindexpr: ident BIND expr { $$ = ast.NewBindExpression($2, $1, $3) }

assignexpr: ident ASSIGN expr { $$ =  ast.NewAssignmentExpression($2, $1, $3) }

expr: bindexpr
	| assignexpr
	| funcexpr
	| ifexpr
	| ident '(' callargs ')' { $$ = ast.NewCallExpression($2, $1, $3) }
	| factor
	| expr '+' factor { $$ = ast.NewInfixExpression($2, $1, $3) }
	| expr '-' factor { $$ = ast.NewInfixExpression($2, $1, $3) }
	| expr comparison expr { $$ = ast.NewInfixExpression($2, $1, $3) }
	;

ifexpr:
	IF '(' expr ')' block { $$ = ast.NewIfExpression($1, $3, $5, nil) }
	| IF '(' expr ')' block ELSE block { $$ = ast.NewIfExpression($1, $3, $5, $7) }
	;

funcexpr:
	FUNCTION '(' funcargs ')' block { $$ = ast.NewFunctionLiteral($1, "", $3, $5) }

funcargs: /*blank*/ { $$ = nil }
	| ident { $$ = []*ast.Identifier{$1.(*ast.Identifier)} }
	| funcargs ',' ident { $$ = append($1, $3.(*ast.Identifier)) }
	;

callargs: /*blank*/ { $$ = nil }
	| expr { $$ = append($$, $1) }
	| callargs ',' expr { $$ = append($$, $3) }
	;

retstmt:
	RETURN expr { $$ = ast.NewReturnStatement($1, $2) }

factor: term
	| factor '*' term { $$ = ast.NewInfixExpression($2, $1, $3) }
	| factor '/' term { $$ = ast.NewInfixExpression($2, $1, $3) }
	| ident
	;

term: number
	| '(' expr ')' { $$ = $2 }
	| '-' term { $$ = ast.NewPrefixExpression($1, $2)}
	;

ident: IDENTIFIER { $$ = ast.NewIdentifier($1, $1.Literal) }
	;

number: NUMBER { $$ = ast.NewIntegerLiteral($1) }
	;

comparison : '>' | '<' | EQ
	;

%%
