package lexer

import "io"

func Ident(l *Lexer) State {

	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			l.EmitIdent()
			return End
		}

		switch nextRune {
		case ';':
			l.EmitIdent()
			return Comment

		case ' ', '\t':
			l.EmitIdent()
			return Start

		case '\n', '.', ',', ':', '(', ')':
			l.EmitIdent()
			l.AddToBuffer(nextRune)
			l.EmitControl()
			return Start

		case '~', '*', '/', '%', '+', '-', '^', '?':
			l.EmitIdent()
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Start

		case '<':
			l.EmitIdent()
			l.AddToBuffer(nextRune)
			return Lt

		case '>':
			l.EmitIdent()
			l.AddToBuffer(nextRune)
			return Gt

		case '=':
			l.EmitIdent()
			l.AddToBuffer(nextRune)
			return Eq

		case '!':
			l.EmitIdent()
			l.AddToBuffer(nextRune)
			return Bang

		case '&':
			l.EmitIdent()
			l.AddToBuffer(nextRune)
			return Amp

		case '|':
			l.EmitIdent()
			l.AddToBuffer(nextRune)
			return Bar

		default:
			l.AddToBuffer(nextRune)
		}
	}
}
