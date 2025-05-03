package language

type ValType int

type Value interface {
	Augment(Value) error
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
	Reg   Reg
}

type RegPointerPreDec struct {
	Value string
	Reg   Reg
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
