package lexer

import "io"

func Register(l *Lexer) State {

	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			l.EmitRegister()
			return End
		}

		switch nextRune {
		case ';':
			l.EmitRegister()
			return Comment

		case ' ', '\t':
			l.EmitRegister()
			return Start

		case '\n', '.', ',', ':', '(', ')':
			l.EmitRegister()
			l.AddToBuffer(nextRune)
			l.EmitControl()
			return Start

		case '~', '*', '/', '%', '+', '-', '^', '?':
			l.EmitRegister()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Start

		case '<':
			l.EmitRegister()
			l.AddToBuffer(nextRune)
			return Lt

		case '>':
			l.EmitRegister()
			l.AddToBuffer(nextRune)
			return Gt

		case '=':
			l.EmitRegister()
			l.AddToBuffer(nextRune)
			return Eq

		case '!':
			l.EmitRegister()
			l.AddToBuffer(nextRune)
			return Bang

		case '&':
			l.EmitRegister()
			l.AddToBuffer(nextRune)
			return Amp

		case '|':
			l.EmitRegister()
			l.AddToBuffer(nextRune)
			return Bar

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.AddToBuffer(nextRune)
			continue

		default:
			l.AddToBuffer(nextRune)
			return Ident
		}
	}
}
