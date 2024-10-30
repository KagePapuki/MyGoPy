package core

import (
	"unicode"
	//"fmt"
	"strings"
	"strconv"
)

/*
0 Keyword: False None True and as assert async await break class continue def del elif else except finally for from global if import in is lambda nonlocal not or pass raise return try while with yield
1 Identifier
2 Literal: number string bytes bool None (中间量: string0' string1" string2''' string3""")
3 Operator: 
	算术运算符：+, -, *, /, //, %, **
	比较运算符：==, !=, <, >, <=, >=
	位运算符：&, |, ^, <<, >>
	赋值运算符：=, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=, **=
	逻辑运算符：and, or, not
4 Separator: ( ) [ ] { } : , .
5 OtherSmbol: COMMENT DECORATOR ELLIPSIS
6 Syntax: INDENT DEDENT NEWLINE(;) START END (中间量: TAB SPACE DOT)
7 Block
*/

type Token struct {
	ttype int
	tcontent string
}

func (t *Token) Ttype() string {
    switch t.ttype {
	case 0:
		return "Keyword"
	case 1:
		return "Identifier"
	case 2:
		return "Literal"
	case 3:
		return "Operator"
	case 4:
		return "Separator"
	case 5:
		return "OtherSmbol"
	case 6:
		return "Syntax"
	case 7:
		return "Block"
	default:
		return ("Unknown(" + strconv.Itoa(t.ttype) + ")")
	}
}

func (t *Token) Tcontent() string {
    return t.tcontent
}

var Command []rune
var I int
var Result []*Token
var Length int

func new_token() *Token {
	var now_token *Token
	var add_token bool = true
	ch := Command[I]
	switch {
	case unicode.IsLetter(ch) || ch == '_':
		now_token = &Token{1, string(ch)}
	case unicode.IsDigit(ch):
		now_token = &Token{2, "number " + string(ch)}
	case ch == '\'':
		if I+2 < Length && Command[I+1] == '\'' && Command[I+2] == '\'' {
			I += 2
			now_token = &Token{2, "string2 "}
		} else {
			now_token = &Token{2, "string0 "}
		}
	case ch == '"':
		if I+2 < Length && Command[I+1] == '"' && Command[I+2] == '"' {
			I += 2
			now_token = &Token{2, "string3 "}
		} else {
			now_token = &Token{2, "string1 "}
		}
	case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '=' || ch == '>' || ch == '<' || ch == '!' || ch == '%' || ch == '&' || ch == '^' || ch == '|':
		now_token = &Token{3, string(ch)}
	case ch == '(' || ch == ')' || ch == '[' || ch == ']' || ch == '{' || ch == '}' || ch == ':' || ch == ',':
		now_token = &Token{4, string(ch)}
	case ch == '#':
		now_token = &Token{5, "COMMENT"}
	case ch == '@':
		now_token = &Token{5, "DECORATOR"}
	case ch == ';' || ch == '\n':
		now_token = &Token{6, "NEWLINE"}
	case ch == '\t':
		now_token = &Token{6, "TAB"}
	case ch == ' ':
		now_token = &Token{6, "SPACE"}
	case ch == '.':
		now_token = &Token{6, "DOT"}
	case ch == '\\' && I+1 < Length && Command[I+1] == '\n':
		//now_token = Result[len(Result) - 1]
		now_token = &Token{8, ""}
		I += 1
		add_token = false
	default:
		now_token = &Token{8, string(ch)}
	}
	if add_token {
		Result = append(Result, now_token)
	}
	return now_token
}

func LexAndYacc(command string) []*Token {
	if command == "" {
		return nil
	}
	Result = []*Token{&Token{6, "START"}}
	var now_token *Token
	//var indent int = 0
	Command = []rune(command)
	Length = len(Command)
	I = 0
	now_token = new_token()
	I += 1

	for I < Length {
		ch := Command[I]
		switch now_token.ttype {
		case 0:
			if unicode.IsLetter(ch) || ch == '_' {
				now_token.tcontent += string(ch)
			} else {
				now_token = new_token()
			}
		case 1:
			if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
				now_token.tcontent += string(ch)
			} else {
				now_token = new_token()
			}
		case 2:
			if strings.HasPrefix(now_token.tcontent, "number ") {
				if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
					now_token.tcontent += string(ch)
				} else {
					now_token = new_token()
				}
			} else if strings.HasPrefix(now_token.tcontent, "string0 ") {
				if ch == '\'' && Command[I-1] != '\\' {
					now_token.tcontent = strings.Replace(now_token.tcontent,"string0 ","string ",1)
				} else if ch == '\\' && Command[I-1] != '\\' && I+1 < Length && Command[I+1] == '\n' {
					I += 1
				} else {
					now_token.tcontent += string(ch)
				}
			} else if strings.HasPrefix(now_token.tcontent, "string1 ") {
				if ch == '"' && Command[I-1] != '\\' {
					now_token.tcontent = strings.Replace(now_token.tcontent,"string1 ","string ",1)
				} else if ch == '\\' && Command[I-1] != '\\' && I+1 < Length && Command[I+1] == '\n' {
					I += 1
				} else {
					now_token.tcontent += string(ch)
				}
			} else if strings.HasPrefix(now_token.tcontent, "string2 ") {
				if I+2 < Length && ch == '\'' && Command[I+1] == '\'' && Command[I+2] == '\'' && Command[I-1] != '\\' {
					now_token.tcontent = strings.Replace(now_token.tcontent,"string2 ","string ",1)
					I += 2
				} else if ch == '\\' && Command[I-1] != '\\' && I+1 < Length && Command[I+1] == '\n' {
					I += 1
				} else {
					now_token.tcontent += string(ch)
				}
			} else if strings.HasPrefix(now_token.tcontent, "string3 ") {
				if I+2 < Length && ch == '"' && Command[I+1] == '"' && Command[I+2] == '"' && Command[I-1] != '\\' {
					now_token.tcontent = strings.Replace(now_token.tcontent,"string3 ","string ",1)
					I += 2
				} else if ch == '\\' && Command[I-1] != '\\' && I+1 < Length && Command[I+1] == '\n' {
					I += 1
				} else {
					now_token.tcontent += string(ch)
				}
			} else if strings.HasPrefix(now_token.tcontent, "string ") {
				now_token = new_token()
			}
		case 3:
			if ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '=' || ch == '>' || ch == '<' || ch == '!' || ch == '%' || ch == '&' || ch == '^' || ch == '|' {
				now_token.tcontent += string(ch)
			} else {
				now_token = new_token()
			}
		case 4:
			now_token = new_token()
		case 5:
			if now_token.tcontent == "COMMENT" {
				if ch == '\n' {
					now_token = new_token()
				}
			} else {
				now_token = new_token()
			}
		case 6:
			now_token = new_token()
		default:
			now_token = new_token()
		}
		I += 1
	}
	
	for I, now_token = range Result {
		switch now_token.ttype {
		case 1:
			if now_token.tcontent == "False" || now_token.tcontent == "None" || now_token.tcontent == "True" || now_token.tcontent == "and" || now_token.tcontent == "as" || now_token.tcontent == "assert" || now_token.tcontent == "async" || now_token.tcontent == "await" || now_token.tcontent == "break" || now_token.tcontent == "class" || now_token.tcontent == "continue" || now_token.tcontent == "def" || now_token.tcontent == "del" || now_token.tcontent == "elif" || now_token.tcontent == "else" || now_token.tcontent == "except" || now_token.tcontent == "finally" || now_token.tcontent == "for" || now_token.tcontent == "from" || now_token.tcontent == "global" || now_token.tcontent == "if" || now_token.tcontent == "import" || now_token.tcontent == "in" || now_token.tcontent == "is" || now_token.tcontent == "lambda" || now_token.tcontent == "nonlocal" || now_token.tcontent == "not" || now_token.tcontent == "or" || now_token.tcontent == "pass" || now_token.tcontent == "raise" || now_token.tcontent == "return" || now_token.tcontent == "try" || now_token.tcontent == "while" || now_token.tcontent == "with" || now_token.tcontent == "yield" {
				now_token.ttype = 0
			}
		}
	}


	Result = append(Result, &Token{6, "END"})
	return Result
}