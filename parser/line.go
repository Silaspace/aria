package parser

type LineType int

const (
	LabelType LineType = 0
	InstrType LineType = 1
	DirType   LineType = 2
	ComType   LineType = 4
	ErrorType LineType = 5
	EOFType   LineType = 6
)

type Line interface {
	Type() LineType
	Fmt() string
}

type EOF struct{}

type Error struct {
	Value string
}

type Comment struct {
	Value string
}

type Label struct {
	Value string
}

type Directive struct {
	Mnemonic string
}

type Instruction struct {
	Mnemonic string
	Op1      Arg
	Op2      Arg
}

func (e *EOF) Type() LineType {
	return EOFType
}

func (e *Error) Type() LineType {
	return ErrorType
}

func (c *Comment) Type() LineType {
	return ComType
}

func (l *Label) Type() LineType {
	return LabelType
}

func (d *Directive) Type() LineType {
	return DirType
}

func (i *Instruction) Type() LineType {
	return InstrType
}
