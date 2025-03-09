package lexer

import "io"

func Comment(l *Lexer) State {
	for {
		nextRune, err := l.GetRune()

		if err == io.EOF {
			l.Emit(TK_COMMENT)
			return End
		}

		switch nextRune {
		case '\n':
			l.Emit(TK_COMMENT)
			return Start

		default:
			l.AddToBuffer(nextRune)
		}
	}
}
