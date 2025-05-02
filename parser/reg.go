package parser

type RegType int
type RegName int
type RegOp int

type Reg interface {
	Type() RegType
	Fmt() string
}

const (
	Single  RegType = 0
	Pair    RegType = 1
	Pointer RegType = 2
	ErrReg  RegType = 3
)

const (
	X RegName = 0
	Y RegName = 1
	Z RegName = 2
)

const (
	None    RegOp = 0
	PostInc RegOp = 1
	PreDec  RegOp = 2
	Disp    RegOp = 3
)

type Register struct {
	Value string
}

type RegPair struct {
	Value string
}

type PointerReg struct {
	Value RegName
	Op    RegOp
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

func (p *PointerReg) Type() RegType {
	return Pointer
}

func (e *RegErr) Type() RegType {
	return ErrReg
}
