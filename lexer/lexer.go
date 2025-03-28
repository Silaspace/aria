package lexer

import (
	"unicode"

	"github.com/silaspace/aria/language"
)

type State func(*Lexer) State

type Reader interface {
	Next() (rune, error)
	Close()
}

type Lexer struct {
	In    Reader
	Out   chan Token
	State State
	Buff  []rune
}

func (l *Lexer) AddToBuffer(r rune) {
	l.Buff = append(l.Buff, r)
}

func (l *Lexer) Close() {
	close(l.Out)
}

func (l *Lexer) DiscardRune() {
	length := len(l.Buff)

	if length > 0 {
		l.Buff = l.Buff[:length-1]
	}
}

func (l *Lexer) Emit(tokentype Type) {
	l.Out <- Token{
		Type:  tokentype,
		Value: string(l.Buff),
	}

	l.Buff = []rune{}
}

func (l *Lexer) EmitControl() {

	// Clear buffer
	control := string(l.Buff)
	l.Buff = []rune{}

	switch control {
	case "\n":
		l.Emit(TK_LINE)
	case ".":
		l.Emit(TK_DOT)
	case ",":
		l.Emit(TK_COMMA)
	case ":":
		l.Emit(TK_COLON)
	case "(":
		l.Emit(TK_LBRAC)
	case ")":
		l.Emit(TK_RBRAC)
	case "=":
		l.Emit(TK_EQ)
	default:
		l.Emit(TK_ERR)
	}
}

func (l *Lexer) EmitIdent() {
	identType := language.Exists(string(l.Buff))

	switch identType {
	case language.INSTR:
		l.Emit(TK_INSTR)

	case language.DIR:
		l.Emit(TK_DIR)

	case language.FUNC:
		l.Emit(TK_FUNC)

	case language.IDENT:
		l.Emit(TK_IDENT)

	default:
		l.Emit(TK_ERR)
	}
}

func (l *Lexer) EmitOperator() {
	_, err := language.GetOp(string(l.Buff))

	if err != nil {
		l.Emit(TK_ERR)
	}

	l.Emit(TK_OP)
}

func (l *Lexer) EmitRegister() {
	l.Buff = l.Buff[1:] // Strip leading r
	l.Emit(TK_REG)
}

func (l *Lexer) GetRune() (rune, error) {
	nextRune, err := l.In.Next()

	if err != nil {
		return ' ', err
	}

	lowerRune := unicode.ToLower(nextRune) // Send all runes to lowercase
	return lowerRune, nil
}

func (l *Lexer) Next() Token {
	for {
		select {
		case nextToken := <-l.Out:
			return nextToken
		default:
			l.State = l.State(l)
		}
	}
}

func (l *Lexer) Reset() {
	l.State = Start
}
