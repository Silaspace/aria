package parser

type ArgType int

type Arg interface {
	Type() ArgType
	Fmt() string
}

const (
	NilArg     ArgType = 0
	RegArg     ArgType = 1
	RegPairArg ArgType = 2
	ExprArg    ArgType = 3
	ErrArg     ArgType = 4
)

type Nil struct{}

type ArgError struct {
	Value string
}

type ArgReg struct {
	Value string
}

type ArgRegPair struct {
	Value string
}

type ArgExpr struct {
	Value Expr
}

func (n *Nil) Type() ArgType {
	return NilArg
}

func (a *ArgError) Type() ArgType {
	return ErrArg
}

func (r *ArgReg) Type() ArgType {
	return RegArg
}

func (r *ArgRegPair) Type() ArgType {
	return RegPairArg
}

func (e *ArgExpr) Type() ArgType {
	return ExprArg
}
