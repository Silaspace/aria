package parser

import (
	"github.com/silaspace/aria/lexer"
)

type Parser struct {
	Lexer  *lexer.Lexer
	Line   uint64
	curtok lexer.Token
}

func (p *Parser) GetCurrentToken() lexer.Token {
	return p.curtok
}

func (p *Parser) GetNextToken() lexer.Token {
	p.curtok = p.Lexer.Next()

	if p.curtok.Type == lexer.TK_LINE {
		p.Line++
	}

	return p.curtok
}

func (p *Parser) Next() Line {
	return ParseLine(p)
}

func (p *Parser) Reset() {
	p.Line = 1
}
