package parser

import (
	"fmt"

	"github.com/silaspace/aria/language"
	"github.com/silaspace/aria/lexer"
)

func ParseReg(p *Parser) Reg {
	token := p.GetCurrentToken()

	switch language.Mnemonic(token.Value) {
	case language.X:
		return ParsePostIncRegPointer(p, X)

	case language.Y:
		return ParsePostIncRegPointer(p, Y)

	case language.Z:
		return ParsePostIncRegPointer(p, Z)
	}

	nextToken := p.GetNextToken()

	switch nextToken.Type {
	case lexer.TK_COLON:
		return ParseRegPair(p)

	default:
		return &Register{
			Value: token.Value,
		}
	}
}

func ParseRegPair(p *Parser) Reg {
	token := p.GetNextToken()

	p.GetNextToken() // Consume

	switch token.Type {
	case lexer.TK_REG:
		return &RegPair{
			Value: token.Value,
		}

	default:
		return &RegErr{
			Value: fmt.Sprintf(
				"Expected register, got '%v'",
				token.Type,
			),
		}

	}
}

func ParsePostIncRegPointer(p *Parser, reg RegName) Reg {
	token := p.GetNextToken()

	switch token.Type {
	case lexer.TK_OP:
		if language.IsAdd(token.Value) {
			nextToken := p.GetNextToken()

			switch nextToken.Type {
			case lexer.TK_IMM:
				p.GetNextToken() // Conusme

				return &PointerReg{
					Value: reg,
					Op:    Disp,
					Disp:  nextToken.Value,
				}

			default:
				return &PointerReg{
					Value: reg,
					Op:    PostInc,
					Disp:  "",
				}
			}
		} else {
			return &RegErr{
				Value: fmt.Sprintf(
					"Unexpected operator %v after register X",
					token.Print(),
				),
			}
		}

	default:
		return &PointerReg{
			Value: reg,
			Op:    None,
			Disp:  "",
		}
	}
}

func ParsePreDecRegPointer(p *Parser) Reg {
	token := p.GetCurrentToken()

	p.GetNextToken() //Consume

	switch language.Mnemonic(token.Value) {
	case language.X:
		return &PointerReg{
			Value: X,
			Op:    PreDec,
			Disp:  "",
		}

	case language.Y:
		return &PointerReg{
			Value: Y,
			Op:    PreDec,
			Disp:  "",
		}

	case language.Z:
		return &PointerReg{
			Value: Z,
			Op:    PreDec,
			Disp:  "",
		}

	default:
		return &RegErr{
			Value: fmt.Sprintf(
				"Unknown register %v in pre-decremnet expression",
				token.Value,
			),
		}
	}
}
