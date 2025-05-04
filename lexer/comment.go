package lexer

import "io"

func Comment(l *Lexer) State {
	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			l.Emit(TK_COM)
			return End
		}

		switch nextRune {
		case '\n':
			l.Emit(TK_COM)
			l.AddToBuffer(nextRune)
			l.EmitControl()
			return Start

		default:
			l.AddToBuffer(nextRune)
		}
	}
}
