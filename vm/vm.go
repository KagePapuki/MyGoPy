package vm

import (
	"github.com/KagePapuki/MyGoPy/decode",
	"os"
)

/*
0 Keyword: False None True and as assert async await break class continue def del elif else except finally for from global if import in is lambda nonlocal not or pass raise return try while with yield
1 Identifier
2 Literal: number string bytes bool None
3 Operator: 
	算术运算符：+, -, *, /, //, %, **
	比较运算符：==, !=, <, >, <=, >=
	位运算符：&, |, ^, <<, >>
	赋值运算符：=, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=, **=
	逻辑运算符：and, or, not
4 Separator: ( ) [ ] { } : ,
5 OtherSmbol: COMMENT DECORATOR ELLIPSIS
6 Syntax: INDENT DEDENT NEWLINE(;) START END DOT
7 Block
*/

type Token struct {
	ttype int
	tcontent string
}

var Path string
var Mem []string
var Address map[string]string
var pointer int

func Init(p string) string {
	_, err := os.Stat(p)
    if err == nil {
		Path = p
		Mem = []string{}
		Address = map[string]string{}
		run([]*Token{&Token{0, "import"},&Token{1, "std"}})
		return "Succeed"
	}
	return "Error"
}

func Run(command []*decode.Token) string {
	pointer = 0
	for pointer < len(command) {
		switch command[pointer].Trawtype() {
		case 0:
			//pass
		case 1:
			//pass
		case 2:
			//pass
		case 3:
			//pass
		case 4:
			//pass
		case 5:
			//pass
		case 6:
			//pass
		case 7:
			//pass
		default:
			//pass
		}
		
		pointer += 1
	}
}

def Get(command []*Token) string {
	//pass
}