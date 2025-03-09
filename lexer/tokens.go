package lexer

import "fmt"

type Type int

type Token struct {
	Type  Type
	Value string
}

const (
	TK_ERR  Type = 10
	TK_EOF  Type = 11
	TK_LINE Type = 12

	TK_COLON Type = 20
	TK_COMMA Type = 21
	TK_DOT   Type = 22
	TK_LBRAC Type = 23
	TK_RBRAC Type = 24
	TK_EQ    Type = 25

	TK_COMMENT Type = 30
	TK_INSTR   Type = 31
	TK_DIR     Type = 32
	TK_FUNC    Type = 33
	TK_IDENT   Type = 34

	TK_OPERATOR Type = 40

	TK_REG Type = 50
	TK_IMM Type = 51
	TK_HEX Type = 52
	TK_OCT Type = 53
	TK_BIN Type = 54
)

func (t *Token) IsEOF() bool {
	return t.Type == TK_EOF
}

func (t *Token) IsErr() bool {
	return t.Type == TK_ERR
}

func (t *Token) Fmt() string {
	return fmt.Sprintf("[%s : %s]\n", _pmap[t.Type], t.Value)
}

func (t *Token) Print() string {
	return _pmap[t.Type]
}

var _pmap = map[Type]string{
	TK_ERR:  "ERR",
	TK_EOF:  "EOF",
	TK_LINE: "LIN",

	TK_COLON: ":  ",
	TK_COMMA: ",  ",
	TK_DOT:   ".  ",
	TK_LBRAC: "(  ",
	TK_RBRAC: ")  ",
	TK_EQ:    "=  ",

	TK_COMMENT:  "COM",
	TK_IDENT:    "IDT",
	TK_INSTR:    "IST",
	TK_FUNC:     "FUN",
	TK_OPERATOR: "OP ",
	TK_REG:      "REG",
	TK_IMM:      "IMM",
	TK_HEX:      "HEX",
	TK_OCT:      "OCT",
	TK_BIN:      "BIN",
}
