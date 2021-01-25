// Code generated by goyacc -o expr.go -p expr expr.y. DO NOT EDIT.

//line expr.y:13

package parser

import __yyfmt__ "fmt"

//line expr.y:14

import (
	"github.com/exfly/newlang/ast"
	"github.com/exfly/newlang/token"
)

//line expr.y:22
type exprSymType struct {
	yys       int
	astTree   ast.Node
	tok       token.Token
	statement ast.Statement
	expr      ast.Expression
	block     *ast.BlockStatement
	id        []*ast.Identifier
	exprs     []ast.Expression
}

const ILLIGAL = 57346
const NUMBER = 57347
const IDENTIFIER = 57348
const BIND = 57349
const ASSIGN = 57350
const FUNCTION = 57351
const EQ = 57352
const IF = 57353
const ELSE = 57354
const WHILE = 57355
const RETURN = 57356
const EOL = 57357

var exprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"';'",
	"','",
	"ILLIGAL",
	"NUMBER",
	"IDENTIFIER",
	"BIND",
	"ASSIGN",
	"FUNCTION",
	"EQ",
	"IF",
	"ELSE",
	"WHILE",
	"RETURN",
	"EOL",
	"'{'",
	"'}'",
	"'>'",
	"'<'",
	"'='",
}

var exprStatenames = [...]string{}

const exprEofCode = 1
const exprErrCode = 2
const exprInitialStackSize = 16

//line expr.y:125

//line yacctab:1
var exprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const exprPrivate = 57344

const exprLast = 120

var exprAct = [...]int{
	62, 3, 4, 2, 21, 10, 63, 68, 19, 16,
	23, 18, 15, 34, 35, 36, 20, 15, 22, 33,
	13, 39, 14, 24, 38, 12, 31, 32, 69, 40,
	45, 43, 43, 37, 46, 47, 49, 48, 52, 19,
	41, 54, 18, 53, 50, 51, 27, 20, 15, 25,
	26, 13, 5, 14, 60, 11, 12, 17, 9, 67,
	61, 65, 58, 30, 59, 64, 8, 66, 21, 70,
	7, 28, 29, 25, 26, 25, 26, 19, 55, 6,
	18, 42, 44, 1, 0, 20, 15, 30, 56, 30,
	57, 19, 0, 0, 18, 28, 29, 28, 29, 20,
	15, 19, 0, 13, 18, 14, 0, 0, 12, 20,
	15, 19, 0, 13, 18, 14, 0, 0, 0, 20,
}

var exprPact = [...]int{
	86, -1000, 86, 0, 71, -1000, -1000, -1000, -1000, -1000,
	11, 7, 96, 25, 16, -1000, -1000, -1000, 96, 106,
	-1000, 0, -1000, -1000, -1000, 72, 72, 96, -1000, -1000,
	-1000, 96, 96, 96, 106, 106, 71, -2, 96, 69,
	-1000, -1000, 7, -1000, 7, 71, 71, 71, 79, 71,
	-1000, -1000, 53, -1000, 45, -1000, -1000, 96, -18, -2,
	-18, 71, -1000, 34, -1000, -13, 3, -1000, -18, -1000,
	-1000,
}

var exprPgo = [...]int{
	0, 83, 79, 70, 66, 2, 5, 58, 55, 9,
	57, 1, 52, 46, 18, 3, 0, 38, 37,
}

var exprR1 = [...]int{
	0, 1, 15, 15, 14, 14, 11, 11, 16, 16,
	2, 3, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 7, 7, 4, 17, 17, 17, 18, 18, 18,
	12, 8, 8, 8, 8, 9, 9, 9, 6, 10,
	13, 13, 13,
}

var exprR2 = [...]int{
	0, 1, 2, 3, 1, 1, 1, 1, 3, 2,
	3, 3, 1, 1, 1, 1, 4, 1, 3, 3,
	3, 5, 7, 5, 0, 1, 3, 0, 1, 3,
	2, 1, 3, 3, 1, 1, 3, 2, 1, 1,
	1, 1, 1,
}

var exprChk = [...]int{
	-1000, -1, -15, -11, -5, -12, -2, -3, -4, -7,
	-6, -8, 22, 17, 19, 14, -9, -10, 8, 5,
	13, -11, -14, 10, 23, 4, 5, -13, 26, 27,
	18, 15, 16, 8, 6, 7, -5, 8, 8, -5,
	-9, -14, -8, -6, -8, -5, -5, -5, -18, -5,
	-9, -9, -17, -6, -5, 9, 9, 11, 9, 11,
	9, -5, -16, 24, -6, -16, -15, 25, 20, 25,
	-16,
}

var exprDef = [...]int{
	0, -2, 1, 0, 6, 7, 12, 13, 14, 15,
	34, 17, 0, 0, 0, 38, 31, 35, 0, 0,
	39, 0, 2, 4, 5, 0, 0, 0, 40, 41,
	42, 0, 0, 27, 0, 0, 30, 24, 0, 0,
	37, 3, 18, 34, 19, 20, 10, 11, 0, 28,
	32, 33, 0, 25, 0, 36, 16, 0, 0, 0,
	0, 29, 23, 0, 26, 21, 0, 9, 0, 8,
	22,
}

var exprTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	8, 9, 6, 4, 11, 5, 3, 7, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 10,
	27, 28, 26, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 24, 3, 25,
}

var exprTok2 = [...]int{
	2, 3, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23,
}

var exprTok3 = [...]int{
	0,
}

var exprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	exprDebug        = 0
	exprErrorVerbose = false
)

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

type exprParser interface {
	Parse(exprLexer) int
	Lookahead() int
}

type exprParserImpl struct {
	lval  exprSymType
	stack [exprInitialStackSize]exprSymType
	char  int
}

func (p *exprParserImpl) Lookahead() int {
	return p.char
}

func exprNewParser() exprParser {
	return &exprParserImpl{}
}

const exprFlag = -1000

func exprTokname(c int) string {
	if c >= 1 && c-1 < len(exprToknames) {
		if exprToknames[c-1] != "" {
			return exprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !exprErrorVerbose {
		return "syntax error"
	}

	for _, e := range exprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + exprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := exprPact[state]
	for tok := TOKSTART; tok-1 < len(exprToknames); tok++ {
		if n := base + tok; n >= 0 && n < exprLast && exprChk[exprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if exprDef[state] == -2 {
		i := 0
		for exprExca[i] != -1 || exprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; exprExca[i] >= 0; i += 2 {
			tok := exprExca[i]
			if tok < TOKSTART || exprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if exprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += exprTokname(tok)
	}
	return res
}

func exprlex1(lex exprLexer, lval *exprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = exprTok1[0]
		goto out
	}
	if char < len(exprTok1) {
		token = exprTok1[char]
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			token = exprTok2[char-exprPrivate]
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		token = exprTok3[i+0]
		if token == char {
			token = exprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = exprTok2[1] /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(token), uint(char))
	}
	return char, token
}

func exprParse(exprlex exprLexer) int {
	return exprNewParser().Parse(exprlex)
}

func (exprrcvr *exprParserImpl) Parse(exprlex exprLexer) int {
	var exprn int
	var exprVAL exprSymType
	var exprDollar []exprSymType
	_ = exprDollar // silence set and not used
	exprS := exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprrcvr.char = -1
	exprtoken := -1 // exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		exprstate = -1
		exprrcvr.char = -1
		exprtoken = -1
	}()
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprtoken), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = exprPact[exprstate]
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprrcvr.char < 0 {
		exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
	}
	exprn += exprtoken
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = exprAct[exprn]
	if exprChk[exprn] == exprtoken { /* valid shift */
		exprrcvr.char = -1
		exprtoken = -1
		exprVAL = exprrcvr.lval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = exprDef[exprstate]
	if exprn == -2 {
		if exprrcvr.char < 0 {
			exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && exprExca[xi+1] == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = exprExca[xi+0]
			if exprn < 0 || exprn == exprtoken {
				break
			}
		}
		exprn = exprExca[xi+1]
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error(exprErrorMessage(exprstate, exprtoken))
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = exprPact[exprS[exprp].yys] + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = exprAct[exprn] /* simulate a shift of "error" */
					if exprChk[exprstate] == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprtoken))
			}
			if exprtoken == exprEofCode {
				goto ret1
			}
			exprrcvr.char = -1
			exprtoken = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= exprR2[exprn]
	// exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if exprp+1 >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = exprR1[exprn]
	exprg := exprPgo[exprn]
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = exprAct[exprg]
	} else {
		exprstate = exprAct[exprj]
		if exprChk[exprstate] != -exprn {
			exprstate = exprAct[exprg]
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:52
		{
			rootAstTree = exprDollar[1].block
		}
	case 2:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:55
		{
			exprVAL.block = ast.NewRowBlockStatement(token.Token{Type: token.LBRACE, Literal: "{"}, nil)
			exprVAL.block.AppendStatement(exprDollar[1].statement)
		}
	case 3:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:56
		{
			exprDollar[1].block.AppendStatement(exprDollar[2].statement)
		}
	case 6:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:61
		{
			exprVAL.statement = ast.NewExpressionStatement(token.Token{Literal: "?"}, exprDollar[1].expr.(ast.Expression))
		}
	case 7:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:62
		{
			exprVAL.statement = exprDollar[1].statement
		}
	case 8:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:65
		{
			exprVAL.block = exprDollar[2].block
		}
	case 9:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:66
		{
			exprVAL.block = ast.NewRowBlockStatement(token.Token{Type: token.LBRACE, Literal: "{"}, nil)
		}
	case 10:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:69
		{
			exprVAL.expr = ast.NewBindExpression(exprDollar[2].tok, exprDollar[1].expr, exprDollar[3].expr)
		}
	case 11:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:71
		{
			exprVAL.expr = ast.NewAssignmentExpression(exprDollar[2].tok, exprDollar[1].expr, exprDollar[3].expr)
		}
	case 16:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line expr.y:77
		{
			exprVAL.expr = ast.NewCallExpression(exprDollar[2].tok, exprDollar[1].expr, exprDollar[3].exprs)
		}
	case 18:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:79
		{
			exprVAL.expr = ast.NewInfixExpression(exprDollar[2].tok, exprDollar[1].expr, exprDollar[3].expr)
		}
	case 19:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:80
		{
			exprVAL.expr = ast.NewInfixExpression(exprDollar[2].tok, exprDollar[1].expr, exprDollar[3].expr)
		}
	case 20:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:81
		{
			exprVAL.expr = ast.NewInfixExpression(exprDollar[2].tok, exprDollar[1].expr, exprDollar[3].expr)
		}
	case 21:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line expr.y:85
		{
			exprVAL.expr = ast.NewIfExpression(exprDollar[1].tok, exprDollar[3].expr, exprDollar[5].block, nil)
		}
	case 22:
		exprDollar = exprS[exprpt-7 : exprpt+1]
//line expr.y:86
		{
			exprVAL.expr = ast.NewIfExpression(exprDollar[1].tok, exprDollar[3].expr, exprDollar[5].block, exprDollar[7].block)
		}
	case 23:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line expr.y:90
		{
			exprVAL.expr = ast.NewFunctionLiteral(exprDollar[1].tok, "", exprDollar[3].id, exprDollar[5].block)
		}
	case 24:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:92
		{
			exprVAL.id = nil
		}
	case 25:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:93
		{
			exprVAL.id = []*ast.Identifier{exprDollar[1].expr.(*ast.Identifier)}
		}
	case 26:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:94
		{
			exprVAL.id = append(exprDollar[1].id, exprDollar[3].expr.(*ast.Identifier))
		}
	case 27:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:97
		{
			exprVAL.exprs = nil
		}
	case 28:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:98
		{
			exprVAL.exprs = append(exprVAL.exprs, exprDollar[1].expr)
		}
	case 29:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:99
		{
			exprVAL.exprs = append(exprVAL.exprs, exprDollar[3].expr)
		}
	case 30:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:103
		{
			exprVAL.statement = ast.NewReturnStatement(exprDollar[1].tok, exprDollar[2].expr)
		}
	case 32:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:106
		{
			exprVAL.expr = ast.NewInfixExpression(exprDollar[2].tok, exprDollar[1].expr, exprDollar[3].expr)
		}
	case 33:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:107
		{
			exprVAL.expr = ast.NewInfixExpression(exprDollar[2].tok, exprDollar[1].expr, exprDollar[3].expr)
		}
	case 36:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:112
		{
			exprVAL.expr = exprDollar[2].expr
		}
	case 37:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:113
		{
			exprVAL.expr = ast.NewPrefixExpression(exprDollar[1].tok, exprDollar[2].expr)
		}
	case 38:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:116
		{
			exprVAL.expr = ast.NewIdentifier(exprDollar[1].tok, exprDollar[1].tok.Literal)
		}
	case 39:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:119
		{
			exprVAL.expr = ast.NewIntegerLiteral(exprDollar[1].tok)
		}
	}
	goto exprstack /* stack new state and value */
}
