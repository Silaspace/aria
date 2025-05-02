package parser

import (
	"fmt"

	"github.com/silaspace/aria/language"
	"github.com/silaspace/aria/lexer"
)

func ParseLine(p *Parser) Line {
	for {
		token := p.GetNextToken()

		switch token.Type {
		case lexer.TK_LINE:
			// Increment line number
			p.Line++
			continue

		case lexer.TK_COMMENT:
			return &Comment{
				Value: token.Value,
				Line:  p.Line,
			}

		case lexer.TK_IDENT:
			return Lab(p)

		case lexer.TK_DOT:
			return Dot(p)

		case lexer.TK_INSTR:
			return Instr(p)

		case lexer.TK_EOF:
			return &EOF{
				Line: p.Line,
			}

		case lexer.TK_ERR:
			return &Error{
				Value: token.Value,
				Line:  p.Line,
			}

		default:
			return &Error{
				Value: fmt.Sprintf(
					"Unexpected token %v after nothing",
					token.Print(),
				),
				Line: p.Line,
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
			Line:  p.Line,
		}

	default:
		return &Error{
			Value: fmt.Sprintf(
				"Unexpected token %v after IDENT",
				token.Print(),
			),
			Line: p.Line,
		}
	}
}

func Dot(p *Parser) Line {
	token := p.GetNextToken()

	switch token.Type {
	case lexer.TK_DIR:
		return Dir(p)

	case lexer.TK_IDENT, lexer.TK_FUNC, lexer.TK_INSTR:
		return &Error{
			Value: fmt.Sprintf(
				"Keyword '%v' is not a directive",
				token.Value,
			),
			Line: p.Line,
		}

	default:
		return &Error{
			Value: fmt.Sprintf(
				"Unexpected token %v after DOT",
				token.Print(),
			),
			Line: p.Line,
		}
	}
}

func Dir(p *Parser) Line {
	token := p.GetCurrentToken()
	mn := language.Mnemonic(token.Value)

	switch mn {
	case language.DIR_DEVICE:
		dirval := DirIdent(p)

		switch dirval := dirval.(type) {
		case *ErrorDirVal:
			return &Error{
				Value: dirval.Value,
				Line:  p.Line,
			}

		default:
			return &Directive{
				Mnemonic: string(mn),
				Value:    dirval,
				Line:     p.Line,
			}
		}

	case language.DIR_EQU:
		dirval := DirAssign(p)

		switch dirval := dirval.(type) {
		case *ErrorDirVal:
			return &Error{
				Value: dirval.Value,
				Line:  p.Line,
			}

		default:
			return &Directive{
				Mnemonic: string(mn),
				Value:    dirval,
				Line:     p.Line - 1,
			}
		}

	default:
		return &Error{
			Value: fmt.Sprintf(
				"Unexpected directive '%v'",
				mn,
			),
		}
	}
}

func Instr(p *Parser) Line {
	token := p.GetCurrentToken()
	p.GetNextToken() // Consume instruction

	arg1 := ParseArg(p)

	if err, ok := arg1.(*ArgError); ok {
		return &Error{
			Value: err.Value,
			Line:  p.Line,
		}
	}

	nextToken := p.GetCurrentToken()

	switch nextToken.Type {
	case lexer.TK_COMMA:
		p.GetNextToken() // Consume ','
		arg2 := ParseArg(p)

		if err, ok := arg2.(*ArgError); ok {
			return &Error{
				Value: err.Value,
				Line:  p.Line,
			}
		}

		thirdToken := p.GetCurrentToken()

		switch thirdToken.Type {
		case lexer.TK_COMMENT, lexer.TK_EOF:
			return &Instruction{
				Mnemonic: token.Value,
				Op1:      arg1,
				Op2:      arg2,
				Line:     p.Line,
			}

		case lexer.TK_LINE:
			// Increment line number
			p.Line++

			return &Instruction{
				Mnemonic: token.Value,
				Op1:      arg1,
				Op2:      arg2,
				Line:     p.Line - 1,
			}

		default:
			return &Error{
				Value: fmt.Sprintf(
					"Unexpected token %v after INST",
					nextToken.Print(),
				),
				Line: p.Line,
			}

		}

	case lexer.TK_COMMENT, lexer.TK_EOF:
		return &Instruction{
			Mnemonic: token.Value,
			Op1:      arg1,
			Op2:      &Nil{},
			Line:     p.Line,
		}

	case lexer.TK_LINE:
		// Increment line number
		p.Line++

		return &Instruction{
			Mnemonic: token.Value,
			Op1:      arg1,
			Op2:      &Nil{},
			Line:     p.Line - 1,
		}

	default:
		return &Error{
			Value: fmt.Sprintf(
				"Unexpected token %v after INST",
				nextToken.Print(),
			),
			Line: p.Line,
		}

	}
}
