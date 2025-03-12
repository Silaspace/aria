package language

import "fmt"

type ValType int

type Value interface {
	Type() ValType
	Fmt() string
}

const (
	NilType    ValType = 0
	ErrType    ValType = 1
	IdentType  ValType = 2
	RegType    ValType = 3
	IntType    ValType = 4
	ListType   ValType = 5
	AssignType ValType = 6
)

type Nil struct{}

type Error struct {
	Value string
}

type Ident struct {
	Value string
}

type Reg struct {
	Value uint64
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

func (r *Reg) Type() ValType {
	return RegType
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

func (n *Nil) Fmt() string {
	return "nil"
}

func (e *Error) Fmt() string {
	return fmt.Sprintf("error (%v)", e.Value)
}

func (i *Ident) Fmt() string {
	return fmt.Sprintf("ident (%v)", i.Value)
}

func (r *Reg) Fmt() string {
	return fmt.Sprintf("reg (%v)", r.Value)
}

func (i *Int) Fmt() string {
	return fmt.Sprintf("int (%v)", i.Value)
}

func (l *List) Fmt() string {
	return fmt.Sprintf("list (%+v)", l.Value)
}

func (a *Assignment) Fmt() string {
	return fmt.Sprintf("assignment (%v = %v)", a.Symbol, a.Value)
}
