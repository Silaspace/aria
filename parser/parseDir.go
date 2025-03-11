package parser

import (
	"fmt"

	"github.com/silaspace/aria/lexer"
)

func DirIdent(p *Parser) DirVal {
	token := p.GetNextToken()

	switch token.Type {
	case lexer.TK_IDENT:
		return &IdentDirVal{
			Value: token.Value,
		}

	default:
		return &ErrorDirVal{
			fmt.Sprintf(
				"Expected ident, got '%v'",
				token.Print(),
			),
		}
	}
}
