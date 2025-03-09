package parser

import (
	"fmt"

	"github.com/silaspace/aria/language"
	"github.com/silaspace/aria/lexer"
)

func GetPrecedence(token lexer.Token) int {
	switch token.Type {
	case lexer.TK_OPERATOR:
		op, _ := language.GetOp(token.Value)
		return op.BindingPower

	case lexer.TK_HEX, lexer.TK_IMM, lexer.TK_OCT, lexer.TK_BIN, lexer.TK_IDENT:
		return 2

	default:
		return 0
	}
}

func ParseExpr(p *Parser, precedence int) Expr {

	left := ParseLeft(p)

	for {
		token := p.GetCurrentToken()
		tbp := GetPrecedence(token)

		if tbp <= precedence {
			break
		}

		left = ParseOp(p, left)
	}

	return left
}

func ParseLeft(p *Parser) Expr {

	token := p.GetCurrentToken()

	switch token.Type {
	case lexer.TK_IDENT:
		val := token.Value
		p.GetNextToken()
		return &Ident{
			Value: val,
		}

	case lexer.TK_HEX:
		val := token.Value
		p.GetNextToken()
		return &Literal{
			Base:  16,
			Value: val,
		}

	case lexer.TK_IMM:
		val := token.Value
		p.GetNextToken()
		return &Literal{
			Base:  10,
			Value: val,
		}

	case lexer.TK_OCT:
		val := token.Value
		p.GetNextToken()
		return &Literal{
			Base:  8,
			Value: val,
		}

	case lexer.TK_BIN:
		val := token.Value
		p.GetNextToken()
		return &Literal{
			Base:  2,
			Value: val,
		}

	case lexer.TK_LBRAC:
		p.GetNextToken() // Consume '('
		expr := ParseExpr(p, 0)
		p.GetNextToken() // Consume ')'
		return expr

	default:
		return &Literal{}
	}
}

func ParseOp(p *Parser, left Expr) Expr {

	token := p.GetCurrentToken()
	tbp := GetPrecedence(token)
	p.GetNextToken()
	right := ParseExpr(p, tbp)

	switch token.Type {
	case lexer.TK_OPERATOR:
		op, _ := language.GetOp(token.Value)
		return &BinopExpr{
			E1: left,
			Op: op,
			E2: right,
		}
	default:
		return &ErrorExpr{
			Value: fmt.Sprintf(
				"Expected operator, got %v",
				token.Print(),
			),
		}
	}
}
