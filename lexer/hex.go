package lexer

import "io"

func Hex(l *Lexer) State {

	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			l.Emit(TK_HEX)
			return End
		}

		switch nextRune {
		case ';':
			l.Emit(TK_HEX)
			return Comment

		case ' ', '\t':
			l.Emit(TK_HEX)
			return Start

		case '\n', '.', ',', ':', '(', ')':
			l.Emit(TK_HEX)
			l.AddToBuffer(nextRune)
			l.EmitControl()
			return Start

		case '~', '*', '/', '%', '+', '-', '^', '?':
			l.Emit(TK_HEX)
			l.AddToBuffer(nextRune)
			l.EmitOperator()
			return Start

		case '<':
			l.Emit(TK_HEX)
			l.AddToBuffer(nextRune)
			return Lt

		case '>':
			l.Emit(TK_HEX)
			l.AddToBuffer(nextRune)
			return Gt

		case '=':
			l.Emit(TK_HEX)
			l.AddToBuffer(nextRune)
			return Eq

		case '!':
			l.Emit(TK_HEX)
			l.AddToBuffer(nextRune)
			return Bang

		case '&':
			l.Emit(TK_HEX)
			l.AddToBuffer(nextRune)
			return Amp

		case '|':
			l.Emit(TK_HEX)
			l.AddToBuffer(nextRune)
			return Bar

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F':
			l.AddToBuffer(nextRune)
			continue

		default:
			l.AddToBuffer(nextRune)
			l.Emit(TK_ERR)
			return Error
		}
	}
}
