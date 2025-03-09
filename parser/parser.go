package parser

import (
	"github.com/silaspace/aria/lexer"
)

type Parser struct {
	Lexer  *lexer.Lexer
	curtok lexer.Token
}

func (p *Parser) GetCurrentToken() lexer.Token {
	return p.curtok
}

func (p *Parser) GetNextToken() lexer.Token {
	p.curtok = p.Lexer.Next()
	return p.curtok
}

func (p *Parser) Next() Line {
	return ParseLine(p)
}
