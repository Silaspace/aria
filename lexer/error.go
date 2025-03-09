package lexer

func Error(l *Lexer) State {
	l.Emit(TK_EOF)
	return Error
}
