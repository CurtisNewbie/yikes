// Code generated by goyacc -o y.go parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

//line parser.y:6
type yySymType struct {
	yys int
	val any
}

const String = 57346
const Number = 57347
const Print = 57348
const Label = 57349
const Type = 57350
const Get = 57351
const Put = 57352
const Post = 57353
const Delete = 57354
const Head = 57355
const Header = 57356
const Body = 57357
const Json = 57358
const JsonStr = 57359
const Comment = 57360
const Bool = 57361
const Write = 57362
const Append = 57363

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"String",
	"Number",
	"Print",
	"Label",
	"Type",
	"Get",
	"Put",
	"Post",
	"Delete",
	"Head",
	"Header",
	"Body",
	"Json",
	"JsonStr",
	"Comment",
	"Bool",
	"Write",
	"Append",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"','",
	"'.'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 5,
	1, 12,
	-2, 48,
	-1, 9,
	1, 7,
	-2, 49,
	-1, 61,
	28, 13,
	-2, 47,
	-1, 62,
	28, 14,
	-2, 46,
	-1, 63,
	28, 15,
	-2, 50,
}

const yyPrivate = 57344

const yyLast = 166

var yyAct = [...]uint8{
	46, 11, 95, 71, 47, 16, 10, 128, 28, 29,
	70, 54, 61, 62, 127, 55, 28, 17, 18, 19,
	20, 21, 125, 27, 29, 51, 124, 63, 28, 45,
	29, 28, 57, 43, 44, 59, 56, 58, 65, 67,
	68, 69, 49, 9, 108, 107, 79, 85, 73, 74,
	75, 76, 106, 77, 83, 25, 24, 14, 5, 15,
	17, 18, 19, 20, 21, 90, 89, 29, 28, 4,
	26, 22, 23, 60, 105, 97, 99, 100, 101, 102,
	96, 104, 103, 32, 33, 34, 35, 134, 111, 114,
	133, 132, 131, 25, 24, 130, 48, 116, 17, 18,
	19, 20, 21, 129, 126, 50, 51, 123, 26, 94,
	93, 92, 91, 88, 87, 34, 35, 115, 86, 42,
	41, 31, 30, 80, 81, 72, 84, 80, 81, 50,
	78, 25, 24, 64, 66, 53, 51, 52, 82, 122,
	51, 112, 82, 121, 113, 109, 26, 120, 110, 119,
	118, 117, 98, 40, 39, 38, 37, 36, 13, 12,
	8, 7, 6, 3, 2, 1,
}

var yyPact = [...]int16{
	51, -1000, -1000, -1000, -1000, 1, -1000, -1000, -1000, -1000,
	-1000, 0, -1000, -1000, 95, 94, 60, 153, 152, 151,
	150, 149, 93, 92, -1000, -1000, -1000, 89, 130, 128,
	8, 126, 127, 127, 127, 127, 111, 111, 111, 111,
	111, 123, 119, 60, -1000, 0, -1000, -1000, -22, -1000,
	91, 87, -1000, -1000, 85, 38, -1000, 37, 84, 83,
	82, -1000, -1000, -1000, 81, 90, -1000, 90, -1000, -1000,
	65, 111, 148, 65, 65, 65, 65, 53, 52, 45,
	-1000, -1000, -1000, 23, 16, 15, 141, 137, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 113, -1000, -1000, -1000,
	-1000, -1000, -1000, 147, 146, 145, 143, 139, 135, 79,
	-2, -6, 76, -14, -21, -1000, -1000, 75, 67, 64,
	63, 62, 59, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 165, 164, 163, 162, 161, 160, 42, 6, 1,
	159, 158, 11, 4, 0, 5, 3, 10, 2,
}

var yyR1 = [...]int8{
	0, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 12, 12, 12, 4, 4, 4, 4,
	4, 4, 4, 10, 10, 10, 11, 11, 11, 6,
	14, 14, 14, 13, 13, 13, 2, 2, 2, 2,
	2, 2, 7, 7, 7, 7, 15, 15, 15, 15,
	15, 16, 17, 17, 17, 18, 18, 18, 8, 8,
	8, 8, 8, 9, 9,
}

var yyR2 = [...]int8{
	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 4, 3, 4,
	4, 4, 4, 6, 6, 6, 6, 6, 6, 4,
	4, 4, 4, 4, 4, 4, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 1, 1, 1,
	1, 2, 1, 0, 2, 0, 2, 2, 4, 4,
	4, 4, 4, 3, 3,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, 18, 7, -4, -5, -6, -7,
	-8, -9, -10, -11, 6, 8, -15, 9, 10, 11,
	12, 13, 20, 21, 5, 4, 19, 22, 30, 30,
	27, 27, 23, 24, 25, 26, 4, 4, 4, 4,
	4, 27, 27, -15, -8, -9, -14, -13, 7, -7,
	16, 17, 7, 7, -12, 7, 28, -9, -8, -13,
	-7, 4, 5, 19, 7, -15, 7, -15, -15, -15,
	-17, -16, 14, -17, -17, -17, -17, -12, 7, -13,
	4, 5, 19, -12, 7, -13, 27, 27, 28, 28,
	28, 28, 28, 28, 28, -18, 15, -16, 4, -18,
	-18, -18, -18, 29, 29, 29, 29, 29, 29, 4,
	7, -9, 4, 7, -9, 4, -14, 4, 4, 4,
	4, 4, 4, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28,
}

var yyDef = [...]int8{
	0, -2, 1, 2, 3, -2, 4, 5, 6, -2,
	8, 9, 10, 11, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 46, 47, 50, 37, 0, 0,
	0, 0, 0, 0, 0, 0, 53, 53, 53, 53,
	53, 0, 0, 36, 38, 39, 40, 41, 48, 49,
	0, 0, 63, 64, 0, 48, 18, 0, 0, 0,
	49, -2, -2, -2, 0, 42, 48, 43, 44, 45,
	55, 52, 0, 55, 55, 55, 55, 0, 0, 0,
	13, 14, 15, 0, 0, 0, 0, 0, 16, 17,
	19, 20, 21, 22, 29, 58, 0, 54, 51, 59,
	60, 61, 62, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 56, 57, 0, 0, 0,
	0, 0, 0, 30, 31, 32, 33, 34, 35, 23,
	24, 25, 26, 27, 28,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	27, 28, 25, 23, 29, 24, 30, 26, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 22,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:48
		{
			PrintYySymDebug(yyDollar[1])
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:54
		{
			PrintYySym(yyDollar[3])
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:55
		{
			PrintGlobalYySym(yyDollar[3])
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:56
		{
			println("")
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:57
		{
			PrintYySym(yyDollar[3])
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:58
		{
			PrintYySym(yyDollar[3])
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:59
		{
			PrintYySym(yyDollar[3])
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:60
		{
			PrintYySym(yyDollar[3])
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:63
		{
			WriteFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:64
		{
			WriteFile(GlobalVarRead(yyDollar[3]), yyDollar[5].val.(string))
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:65
		{
			WriteFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:68
		{
			AppendFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:69
		{
			AppendFile(GlobalVarRead(yyDollar[3]), yyDollar[5].val.(string))
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:70
		{
			AppendFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:73
		{
			PrintType(yyDollar[3])
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:76
		{
			yyVAL = yySymType{val: StrToMap(yyDollar[3].val)}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:77
		{
			yyVAL = yySymType{val: StrToMap(GlobalVarRead(yyDollar[3]))}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:78
		{
			yyVAL = yySymType{val: StrToMap(yyDollar[3].val)}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:81
		{
			yyVAL = yySymType{val: ToJsonStr(yyDollar[3].val)}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:82
		{
			yyVAL = yySymType{val: ToJsonStr(GlobalVarRead(yyDollar[3]))}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:83
		{
			yyVAL = yySymType{val: ToJsonStr(yyDollar[3].val)}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:86
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:87
		{
			SyntaxError()
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:88
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:89
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:90
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:91
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:94
		{
			yyVAL = yySymType{val: ValAdd(yyDollar[1].val, yyDollar[3].val)}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:95
		{
			yyVAL = yySymType{val: ValMinus(yyDollar[1].val, yyDollar[3].val)}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:96
		{
			yyVAL = yySymType{val: ValMul(yyDollar[1].val, yyDollar[3].val)}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:97
		{
			yyVAL = yySymType{val: ValDiv(yyDollar[1].val, yyDollar[3].val)}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:100
		{
			yyVAL = yyDollar[1]
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:101
		{
			yyVAL = yyDollar[1]
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:102
		{
			yyVAL = yySymType{val: GlobalVarRead(yyDollar[1])}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:103
		{
			yyVAL = yyDollar[1]
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:104
		{
			yyVAL = yyDollar[1]
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:107
		{
			yyVAL = yyDollar[2]
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:110
		{
			yyVAL = yyDollar[1]
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:112
		{
			yyVAL = joinHeaders(yyDollar[1], yyDollar[2])
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:116
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:117
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:120
		{
			yyVAL = HttpSend("GET", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:121
		{
			yyVAL = HttpSend("PUT", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:122
		{
			yyVAL = HttpSend("POST", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:123
		{
			yyVAL = HttpSend("DELETE", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:124
		{
			yyVAL = HttpSend("HEAD", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:131
		{
			v := GlobalVarRead(yyDollar[1])
			yyVAL = yySymType{val: WalkField(v, yyDollar[3].val)}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:135
		{
			yyVAL = yySymType{val: WalkField(yyDollar[1].val, yyDollar[3].val)}
		}
	}
	goto yystack /* stack new state and value */
}
