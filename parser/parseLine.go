package parser

import (
	"fmt"

	"github.com/silaspace/aria/lexer"
)

func ParseLine(p *Parser) Line {
	for {
		token := p.GetNextToken()

		switch token.Type {
		case lexer.TK_LINE:
			continue

		case lexer.TK_COMMENT:
			return &Comment{
				Value: token.Value,
			}

		case lexer.TK_IDENT:
			return Lab(p)

		case lexer.TK_DOT:
			return Dot(p)

		case lexer.TK_INSTR:
			return Instr(p)

		case lexer.TK_EOF:
			return &EOF{}

		case lexer.TK_ERR:
			return &Error{
				Value: token.Value,
			}

		default:
			return &Error{
				fmt.Sprintf(
					"Unexpected token %v after nothing",
					token.Print(),
				),
			}
		}
	}
}

func Lab(p *Parser) Line {
	ident := p.GetCurrentToken().Value
	token := p.GetNextToken()

	switch token.Type {
	case lexer.TK_COLON:
		return &Label{
			Value: ident,
		}

	case lexer.TK_DIR, lexer.TK_INSTR:
		return &Error{
			Value: fmt.Sprintf(
				"Keyword '%v' cannot be used as label",
				token.Value,
			),
		}

	default:
		return &Error{
			fmt.Sprintf(
				"Unexpected token %v after IDENT",
				token.Print(),
			),
		}
	}
}

func Dot(p *Parser) Line {
	token := p.GetNextToken()

	switch token.Type {
	case lexer.TK_DIR:
		return Dir(p)

	case lexer.TK_IDENT:
		return &Error{
			fmt.Sprintf(
				"Identifier '%v' is not a directive",
				token.Value,
			),
		}

	default:
		return &Error{
			fmt.Sprintf(
				"Unexpected token %v after DOT",
				token.Print(),
			),
		}
	}
}

func Dir(p *Parser) Line {
	return &Error{
		Value: "Dir not implemented",
	}
}

func Instr(p *Parser) Line {
	token := p.GetCurrentToken()
	p.GetNextToken() // Consume instruction

	arg1 := ParseArg(p)

	if err, ok := arg1.(*ArgError); ok {
		return &Error{
			Value: err.Value,
		}
	}

	nextToken := p.GetCurrentToken()

	switch nextToken.Type {
	case lexer.TK_COMMA:
		p.GetNextToken() // Consume ','
		arg2 := ParseArg(p)

		switch arg2 := arg2.(type) {
		case *ArgError:
			return &Error{
				Value: arg2.Value,
			}

		default:
			return &Instruction{
				Mnemonic: token.Value,
				Op1:      arg1,
				Op2:      arg2,
			}
		}

	case lexer.TK_LINE, lexer.TK_COMMENT, lexer.TK_EOF:
		return &Instruction{
			Mnemonic: token.Value,
			Op1:      arg1,
			Op2:      &Nil{},
		}

	default:
		return &Error{
			fmt.Sprintf(
				"Unexpected token %v after IST",
				nextToken.Print(),
			),
		}

	}
}
