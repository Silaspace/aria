package parser

import (
	"fmt"

	"github.com/silaspace/aria/language"
	"github.com/silaspace/aria/lexer"
)

func ParseArg(p *Parser) Arg {
	token := p.GetCurrentToken()

	switch token.Type {
	case lexer.TK_EOF, lexer.TK_COM:
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
		return ParseExprArg(p)
	}
}

func ParseExprArg(p *Parser) Arg {
	token := p.GetCurrentToken()

	if language.IsSub(token.Value) {
		nextToken := p.GetNextToken()

		switch nextToken.Type {
		case lexer.TK_REG:
			r := ParsePreDecRegPointer(p)
			return &ArgReg{
				Value: r,
			}

		// Carry on parsing expr in the form -(e)
		default:
			op, _ := language.GetOp(token.Value)
			tbp := GetPrecedence(token, true)
			p.GetNextToken()
			expr := ParseExpr(p, tbp)

			e := &MonopExpr{
				E1:     expr,
				Symbol: token.Value,
				Op:     op,
			}

			return &ArgExpr{
				Value: e,
			}
		}

	} else {
		e := ParseExpr(p, 0)
		return &ArgExpr{
			Value: e,
		}
	}

}
