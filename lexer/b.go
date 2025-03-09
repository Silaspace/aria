package lexer

import "io"

func B(l *Lexer) State {

	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			l.EmitIdent()
			return End
		}

		switch nextRune {
		case '0', '1':
			l.AddToBuffer(nextRune)
			return Binary

		default:
			l.AddToBuffer(nextRune)
			l.Emit(TK_ERR)
			return Error
		}
	}
}
