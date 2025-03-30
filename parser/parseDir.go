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

func DirAssign(p *Parser) DirVal {
	token := p.GetNextToken()

	switch token.Type {
	case lexer.TK_IDENT:
		nextToken := p.GetNextToken()

		switch nextToken.Type {
		case lexer.TK_EQ:
			p.GetNextToken() // Consume =
			expr := ParseExpr(p, 0)

			return &AssignDirVal{
				Symbol: token.Value,
				Value:  expr,
			}

		default:
			return &ErrorDirVal{
				fmt.Sprintf(
					"Expected =, got '%v'",
					token.Print(),
				),
			}
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
