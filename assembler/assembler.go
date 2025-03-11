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
	Device  device.Device
	Parser  parser.Parser
	PC      uint64
	Reader  Reader
	Symbols map[string]uint64
	Writer  Writer
}

func (a *Assembler) AddSymbol(symbol string, value uint64) {
	a.Symbols[symbol] = value
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

	for {
		line := a.Parser.Next()

		if _, ok := line.(*parser.EOF); ok {
			break
		}

		switch line := line.(type) {
		case *parser.Comment:
			continue

		case *parser.Directive:
			dir, err := language.GetDir(line.Mnemonic)

			if err != nil {
				return err
			}

			val := EvalDirVal(line.Value, a.Symbols)
			dir.Execute(a, val)

		case *parser.Label:
			_, exists := a.Symbols[line.Value]

			if exists {
				return a.error("duplicate label %v", line.Value)
			}

			a.AddSymbol(line.Value, a.PC)

		case *parser.Instruction:
			instr, err := language.GetInstr(line.Mnemonic, &a.Device)

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
			instr, err := language.GetInstr(line.Mnemonic, &a.Device)

			if err != nil {
				return err
			}

			relative := instr.IsRelative()

			if line.Op1.Type() != parser.NilArg {
				operand, err := EvalArg(line.Op1, a.Symbols, relative, a.PC)

				if err != nil {
					return err
				}

				err = instr.Apply1(operand)

				if err != nil {
					return err
				}
			}

			if line.Op2.Type() != parser.NilArg {
				operand, err := EvalArg(line.Op2, a.Symbols, relative, a.PC)

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

func (a *Assembler) SetDevice(name string) {
	dev, err := device.NewDevice(name)

	if err != nil {
		a.error(err.Error())
	}

	a.Device = *dev
}

func (a *Assembler) error(fstr string, args ...interface{}) error {
	err := fmt.Errorf(fstr, args...)
	return err
}
