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
	Fmt() string
	Number() uint64
	Type() LineType
}

type EOF struct {
	Line uint64
}

type Error struct {
	Value string
	Line  uint64
}

type Comment struct {
	Value string
	Line  uint64
}

type Label struct {
	Value string
	Line  uint64
}

type Directive struct {
	Mnemonic string
	Value    DirVal
	Line     uint64
}

type Instruction struct {
	Mnemonic string
	Op1      Arg
	Op2      Arg
	Line     uint64
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

func (e *EOF) Number() uint64 {
	return e.Line
}

func (e *Error) Number() uint64 {
	return e.Line
}

func (c *Comment) Number() uint64 {
	return c.Line
}

func (l *Label) Number() uint64 {
	return l.Line
}

func (d *Directive) Number() uint64 {
	return d.Line
}

func (i *Instruction) Number() uint64 {
	return i.Line
}
