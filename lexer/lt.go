package lexer

import "io"

func Lt(l *Lexer) State {
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

		case '<', '=':
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Start

		case '>':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Gt

		case '!':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Bang

		case '&':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Amp

		case '|':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Bar

		case 'r':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return R

		case '0':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Zero

		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Imm

		default:
			l.EmitOperator()
			l.AddToBuffer(nextRune)
			return Ident
		}
	}
}
