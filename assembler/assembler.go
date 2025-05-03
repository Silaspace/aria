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
	Line    uint64
	Parser  parser.Parser
	PC      uint64
	Reader  Reader
	Symbols map[string]uint64
	Writer  Writer
}

func (a *Assembler) AddSymbol(symbol string, value uint64) error {
	_, exists := a.Symbols[symbol]

	if exists {
		return a.error("duplicate label %v", symbol)
	}

	a.Symbols[symbol] = value
	return nil
}

func (a *Assembler) Close() {
	a.Reader.Close()
	a.Writer.Close()
}

func (a *Assembler) HardReset() error {
	err := a.SoftReset()

	if err != nil {
		return err
	}

	a.Symbols = map[string]uint64{}
	a.Device = *device.DefaultDevice()

	return nil
}

func (a *Assembler) GetNextLine() parser.Line {
	line := a.Parser.Next()
	a.Line = line.Number()
	return line
}

func (a *Assembler) Run() error {
	/*
		PASS 1 - Record lables and directives
	*/

	err := a.HardReset()

	if err != nil {
		return err
	}

	for {
		line := a.GetNextLine()

		if _, ok := line.(*parser.EOF); ok {
			break
		}

		switch line := line.(type) {
		case *parser.Comment:
			continue

		case *parser.Directive:
			dir, err := language.GetDir(line.Mnemonic)

			if err != nil {
				return a.wrap(err)
			}

			val := EvalDirVal(line.Value, a.Symbols)
			dir.Execute(a, val)

		case *parser.Label:
			err := a.AddSymbol(line.Value, a.PC)

			if err != nil {
				return a.wrap(err)
			}

		case *parser.Instruction:
			instr, err := language.GetInstr(line.Mnemonic, &a.Device)

			if err != nil {
				return a.wrap(err)
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

	err = a.SoftReset()

	if err != nil {
		return err
	}

	for {
		line := a.GetNextLine()

		if _, ok := line.(*parser.EOF); ok {
			break
		}

		switch line := line.(type) {
		case *parser.Comment, *parser.Directive, *parser.Label:
			continue

		case *parser.Instruction:
			instr, err := language.GetInstr(line.Mnemonic, &a.Device)

			if err != nil {
				return a.wrap(err)
			}

			relative := instr.IsRelative()

			op1 := EvalArg(line.Op1, a.Symbols, relative, a.PC)
			op2 := EvalArg(line.Op2, a.Symbols, relative, a.PC)

			// Check arguments against one another for undefined behaviour
			if err := op1.Augment(op2); err != nil {
				return a.wrap(err)
			}

			if err := op2.Augment(op1); err != nil {
				return a.wrap(err)
			}

			if err := instr.Apply1(op1); err != nil {
				return a.wrap(err)
			}

			if err := instr.Apply2(op2); err != nil {
				return a.wrap(err)
			}

			if instr.IsLong() {
				a.PC += 2
			} else {
				a.PC += 1
			}

			bytes := instr.Encode()
			err = a.Writer.Write(bytes)

			if err != nil {
				return a.wrap(err)
			}

		case *parser.Error:
			return a.error(line.Value)

		default:
			return a.error("unknown parsing error")
		}
	}

	return nil
}

func (a *Assembler) SetDevice(name string) error {
	dev, err := device.NewDevice(name)

	if err != nil {
		return err
	}

	a.Device = *dev
	return nil
}

func (a *Assembler) SoftReset() error {
	err := a.Reader.Reset()

	if err != nil {
		return err
	}

	a.Parser.Reset()
	a.Parser.Lexer.Reset()
	a.PC = 0
	return nil
}

func (a *Assembler) error(fstr string, args ...interface{}) error {
	msg := fmt.Sprintf(fstr, args...)
	err := fmt.Errorf("%v on line %v", msg, a.Line)
	return err
}

func (a *Assembler) wrap(err error) error {
	msg := err.Error()
	return a.error(msg)
}
