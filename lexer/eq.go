package lexer

import "io"

func Eq(l *Lexer) State {
	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			return End
		}

		switch nextRune {
		case ';':
			l.EmitControl()
			return Comment

		case ' ', '\t':
			l.EmitControl()
			return Start

		case '\n', '.', ',', ':', '(', ')':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			l.EmitControl()
			return Start

		case '~', '*', '/', '%', '+', '-', '^', '?':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Start

		case '=':
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Start

		case '>':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Gt

		case '<':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Lt

		case '!':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Bang

		case '&':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Amp

		case '|':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Bar

		case 'r':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			return R

		case '0':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			return Zero

		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.EmitControl()
			l.AddToBuffer(nextRune)
			return Imm

		default:
			l.EmitControl()
			l.AddToBuffer(nextRune)
			return Ident
		}
	}
}
