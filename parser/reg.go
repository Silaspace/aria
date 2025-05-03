package parser

type RegType int
type RegOp int

type Reg interface {
	Type() RegType
	Fmt() string
}

const (
	Single         RegType = 0
	Pair           RegType = 1
	Pointer        RegType = 2
	PointerPostInc RegType = 2
	PointerPreDec  RegType = 2
	PointerDisp    RegType = 2
	ErrReg         RegType = 3
)

type Register struct {
	Value string
}

type RegPair struct {
	Value string
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
	Disp  string
}

type RegErr struct {
	Value string
}

func (r *Register) Type() RegType {
	return Single
}

func (r *RegPair) Type() RegType {
	return Pair
}

func (p *RegPointer) Type() RegType {
	return Pointer
}

func (p *RegPointerPostInc) Type() RegType {
	return PointerPostInc
}

func (p *RegPointerPreDec) Type() RegType {
	return PointerPreDec
}

func (p *RegPointerDisp) Type() RegType {
	return PointerDisp
}

func (e *RegErr) Type() RegType {
	return ErrReg
}
