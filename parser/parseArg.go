package parser

import (
	"fmt"

	"github.com/silaspace/aria/language"
	"github.com/silaspace/aria/lexer"
)

func ParseArg(p *Parser) Arg {
	token := p.GetCurrentToken()

	switch token.Type {
	case lexer.TK_EOF, lexer.TK_COMMENT:
		return &Nil{}

	case lexer.TK_LINE:
		// Increment line number
		p.Line++
		return &Nil{}

	case lexer.TK_REG:
		// Special case of the program counter
		if token.Value == string(language.PC) {
			e := ParseExpr(p, 0)
			return &ArgExpr{
				Value: e,
			}
		}

		nextToken := p.GetNextToken()

		switch nextToken.Type {
		case lexer.TK_COLON:
			thirdToken := p.GetNextToken()

			p.GetNextToken() // Consume

			switch thirdToken.Type {
			case lexer.TK_REG:
				return &ArgRegPair{
					Value: thirdToken.Value,
				}

			default:
				return &ArgError{
					Value: fmt.Sprintf(
						"Expected register, got '%v'",
						thirdToken.Type,
					),
				}

			}

		default:
			return &ArgReg{
				Value: token.Value,
			}
		}

	case lexer.TK_INSTR, lexer.TK_DIR:
		return &ArgError{
			Value: fmt.Sprintf(
				"Keyword '%v' cannot be operand",
				token.Value,
			),
		}

	default:
		e := ParseExpr(p, 0)
		return &ArgExpr{
			Value: e,
		}
	}
}
