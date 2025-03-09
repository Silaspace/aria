package lexer

const BUFFER_LEN int = 30

func NewLexer(in Reader) *Lexer {
	return &Lexer{
		In:    in,
		Out:   make(chan Token, BUFFER_LEN),
		State: Start,
	}
}
