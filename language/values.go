package language

import (
	"fmt"
)

type ValType int
type PointerOp uint64

type Value interface {
	Type() ValType
	Fmt() string
}

const (
	NilType               ValType = 0
	ErrType               ValType = 1
	IdentType             ValType = 2
	RegType               ValType = 3
	RegPairType           ValType = 4
	RegPointerType        ValType = 5
	RegPointerPostIncType ValType = 6
	RegPointerPreDecType  ValType = 7
	RegPointerDispType    ValType = 8
	IntType               ValType = 9
	ListType              ValType = 10
	AssignType            ValType = 11
)

const (
	None    PointerOp = 0
	PostInc PointerOp = 1
	PreDec  PointerOp = 2
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

type RegPair struct {
	Value uint64
}

type RegPointer struct {
	Value string
}

type RegPointerPostInc struct {
	Value string
}

type RegPointerPreDec struct {
	Value string
}

type RegPointerDisp struct {
	Value string
	Disp  uint64
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

func (r *RegPair) Type() ValType {
	return RegPairType
}

func (r *RegPointer) Type() ValType {
	return RegPointerType
}

func (r *RegPointerPostInc) Type() ValType {
	return RegPointerPostIncType
}

func (r *RegPointerPreDec) Type() ValType {
	return RegPointerPreDecType
}

func (r *RegPointerDisp) Type() ValType {
	return RegPointerDispType
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

func (r *RegPair) Fmt() string {
	return fmt.Sprintf("reg (%v : %v)", r.Value+1, r.Value)
}

func (r *RegPointer) Fmt() string {
	return fmt.Sprintf("reg (%v)", r.Value)
}

func (r *RegPointerPostInc) Fmt() string {
	return fmt.Sprintf("reg (%v+)", r.Value)
}

func (r *RegPointerPreDec) Fmt() string {
	return fmt.Sprintf("reg (-%v)", r.Value)
}

func (r *RegPointerDisp) Fmt() string {
	return fmt.Sprintf("reg (%v+%v)", r.Value, r.Disp)
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
