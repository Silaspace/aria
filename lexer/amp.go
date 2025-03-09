package lexer

import "io"

func Amp(l *Lexer) State {
	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			return End
		}

		switch nextRune {
		case ';':
			l.EmitOperator()
			return Comment

		case ' ', '\t':
			l.EmitOperator()
			return Start

		case '\n', '.', ',', ':', '(', ')':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			l.EmitControl()
			return Start

		case '~', '*', '/', '%', '+', '-', '^', '?':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Start

		case '>':
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Gt

		case '<':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Lt

		case '=':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Eq

		case '!':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Bang

		case '&':
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Start

		case '|':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Bar

		case 'r':
			l.AddToBuffer(nextRune)
			return R

		case '0':
			l.AddToBuffer(nextRune)
			return Zero

		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.AddToBuffer(nextRune)
			return Imm

		default:
			l.AddToBuffer(nextRune)
			return Ident
		}
	}
}
