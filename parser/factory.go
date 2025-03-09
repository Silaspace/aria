package parser

import "github.com/silaspace/aria/lexer"

const BUFFER_LEN int = 30

func NewParser(in *lexer.Lexer) *Parser {
	return &Parser{
		Lexer: in,
	}
}
