package language

import "fmt"

type ValType int

type Value interface {
	Type() ValType
}

const (
	NilType    ValType = 0
	ErrType    ValType = 1
	IdentType  ValType = 2
	IntType    ValType = 3
	ListType   ValType = 4
	AssignType ValType = 5
)

type Nil struct{}

type Error struct {
	Value string
}

type Ident struct {
	Value string
}

type Int struct {
	Value uint64
}

type List struct {
	Value []uint64
}

type Assignment struct {
	Symbol string
	Value  uint64
}

func (n *Nil) Type() ValType {
	return NilType
}

func (e *Error) Type() ValType {
	return ErrType
}

func (i *Ident) Type() ValType {
	return IdentType
}

func (i *Int) Type() ValType {
	return IntType
}

func (l *List) Type() ValType {
	return ListType
}

func (a *Assignment) Type() ValType {
	return AssignType
}

type Directive struct {
	Execute func(Assembler, Value) error
}

type Assembler interface {
	AddSymbol(string, uint64)
	SetDevice(string)
}

const (
	DIR_DEVICE Mnemonic = "device"
	DIR_EQU    Mnemonic = "equ"
)

var Directives = map[Mnemonic]Directive{
	DIR_DEVICE: {
		Execute: func(a Assembler, v Value) error {
			switch v := v.(type) {
			case *Ident:
				a.SetDevice(v.Value)
				return nil

			default:
				return fmt.Errorf("expected device, got '%v'", v)
			}
		},
	},
	DIR_EQU: {
		Execute: func(a Assembler, v Value) error {
			switch v := v.(type) {
			case *Assignment:
				a.AddSymbol(v.Symbol, v.Value)
				return nil

			default:
				return fmt.Errorf("expected assignment, got '%v'", v)
			}
		},
	},
}
