package lexer

func End(l *Lexer) State {
	l.Emit(TK_EOF)
	return Error
}
