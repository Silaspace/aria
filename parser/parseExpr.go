package parser

import (
	"fmt"

	"github.com/silaspace/aria/language"
	"github.com/silaspace/aria/lexer"
)

func GetPrecedence(token lexer.Token, unary bool) int {
	switch token.Type {
	case lexer.TK_OP:
		op, _ := language.GetOp(token.Value)

		// Hardcode unary minus precedence
		if unary && token.Value == string(language.SUB) {
			return 14
		}

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
		tbp := GetPrecedence(token, false)

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

	case lexer.TK_FUNC:
		function, _ := language.GetFunc(token.Value)
		p.GetNextToken()

		p.GetNextToken() // Consume '('
		expr := ParseExpr(p, 0)
		p.GetNextToken() // Consume ')'

		return &FuncExpr{
			E1:     expr,
			Symbol: token.Value,
			Func:   function,
		}

	case lexer.TK_OP:
		op, _ := language.GetOp(token.Value)

		if op.IsUnary() {
			tbp := GetPrecedence(token, true)
			p.GetNextToken()
			expr := ParseExpr(p, tbp)

			return &MonopExpr{
				E1:     expr,
				Symbol: token.Value,
				Op:     op,
			}
		} else {
			return &ErrorExpr{
				Value: fmt.Sprintf(
					"Binary operator %v with no left expression",
					token.Print(),
				),
			}
		}

	default:
		return &ErrorExpr{
			Value: fmt.Sprintf(
				"Unexpected token %v in expression",
				token.Print(),
			),
		}
	}
}

func ParseOp(p *Parser, left Expr) Expr {

	token := p.GetCurrentToken()
	tbp := GetPrecedence(token, false)
	p.GetNextToken()
	right := ParseExpr(p, tbp)

	switch token.Type {
	case lexer.TK_OP:
		op, _ := language.GetOp(token.Value)

		if op.IsBinary() {
			return &BinopExpr{
				E1:     left,
				E2:     right,
				Symbol: token.Value,
				Op:     op,
			}
		} else {
			return &ErrorExpr{
				Value: fmt.Sprintf(
					"Unary operator %v with left expression supplied",
					token.Print(),
				),
			}
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
