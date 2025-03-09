package assembler

import (
	"fmt"

	"github.com/silaspace/aria/device"
	"github.com/silaspace/aria/language"
	"github.com/silaspace/aria/parser"
)

type Reader interface {
	Next() (rune, error)
	Reset() error
	Close()
}

type Writer interface {
	Write([]byte) error
	Close()
}

type Assembler struct {
	Reader Reader
	Writer Writer
	Parser parser.Parser
	PC     uint64
}

func (a *Assembler) Close() {
	a.Reader.Close()
	a.Writer.Close()
}

func (a *Assembler) Reset() error {
	err := a.Reader.Reset()
	if err != nil {
		return err
	}
	a.Parser.Lexer.Reset()
	a.PC = 0
	return nil
}

func (a *Assembler) Run() error {
	/*
		PASS 1 - Record lables and directives
	*/

	err := a.Reset()

	if err != nil {
		return err
	}

	symbols := map[string]uint64{}
	dev := device.DefaultDevice()

	for {
		line := a.Parser.Next()

		if _, ok := line.(*parser.EOF); ok {
			break
		}

		switch line := line.(type) {
		case *parser.Comment:
			continue

		case *parser.Directive:
			continue

		case *parser.Label:
			_, exists := symbols[line.Value]

			if exists {
				return a.error("duplicate label %v", line.Value)
			}

			symbols[line.Value] = a.PC

		case *parser.Instruction:
			instr, err := language.GetInstr(line.Mnemonic, dev)

			if err != nil {
				return err
			}

			if instr.IsLong() {
				a.PC += 2
			} else {
				a.PC += 1
			}

		case *parser.Error:
			return a.error(line.Value)

		default:
			return a.error("unknown parsing error")
		}
	}

	/*
		PASS 2 - Generate code
	*/

	err = a.Reset()

	if err != nil {
		return err
	}

	for {

		a.Parser.Lexer.Reset()

		line := a.Parser.Next()

		if _, ok := line.(*parser.EOF); ok {
			break
		}

		switch line := line.(type) {
		case *parser.Comment, *parser.Directive, *parser.Label:
			continue

		case *parser.Instruction:
			instr, err := language.GetInstr(line.Mnemonic, dev)

			if err != nil {
				return err
			}

			relative := instr.IsRelative()

			if line.Op1.Type() != parser.NilArg {
				operand, err := EvalArg(line.Op1, symbols, relative, a.PC)

				if err != nil {
					return err
				}

				err = instr.Apply1(operand)

				if err != nil {
					return err
				}
			}

			if line.Op2.Type() != parser.NilArg {
				operand, err := EvalArg(line.Op2, symbols, relative, a.PC)

				if err != nil {
					return err
				}

				err = instr.Apply2(operand)

				if err != nil {
					return err
				}
			}

			if instr.IsLong() {
				a.PC += 2
			} else {
				a.PC += 1
			}

			bytes := instr.Encode()
			err = a.Writer.Write(bytes)

			if err != nil {
				return err
			}

		case *parser.Error:
			return a.error(line.Value)

		default:
			return a.error("unknown parsing error")
		}
	}

	return nil
}

func (a *Assembler) error(fstr string, args ...interface{}) error {
	//msg := fmt.Sprintf(fstr, args...)
	err := fmt.Errorf(fstr, args...)
	return err
}
