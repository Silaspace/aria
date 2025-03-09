package lexer

import "io"

func Start(l *Lexer) State {
	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			return End
		}

		switch nextRune {
		case ';':
			return Comment

		case ' ', '\t':
			continue

		case '\n', '.', ',', ':', '(', ')':
			l.AddToBuffer(nextRune)
			l.EmitControl()
			continue

		case '~', '*', '/', '%', '+', '-', '^', '?':
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			continue

		case '<':
			l.AddToBuffer(nextRune)
			return Lt

		case '>':
			l.AddToBuffer(nextRune)
			return Gt

		case '=':
			l.AddToBuffer(nextRune)
			return Eq

		case '!':
			l.AddToBuffer(nextRune)
			return Bang

		case '&':
			l.AddToBuffer(nextRune)
			return Amp

		case '|':
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
