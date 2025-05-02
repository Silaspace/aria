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
		// Special case of the program counter, should return Expr type
		if language.IsPC(token.Value) {
			e := ParseExpr(p, 0)
			return &ArgExpr{
				Value: e,
			}
		}

		r := ParseReg(p)
		return &ArgReg{
			Value: r,
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
