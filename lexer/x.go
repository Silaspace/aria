package lexer

import "io"

func X(l *Lexer) State {

	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			l.EmitIdent()
			return End
		}

		switch nextRune {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F':
			l.AddToBuffer(nextRune)
			return Hex

		default:
			l.AddToBuffer(nextRune)
			l.Emit(TK_ERR)
			return Error
		}
	}
}
