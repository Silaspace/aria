package parser

type ArgType int

type Arg interface {
	Type() ArgType
	Fmt() string
}

const (
	NilArg  ArgType = 0
	RegArg  ArgType = 1
	ExprArg ArgType = 2
	ErrArg  ArgType = 3
)

type Nil struct{}

type ArgError struct {
	Value string
}

type ArgReg struct {
	Value string
	E1    Expr
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

func (e *ArgExpr) Type() ArgType {
	return ExprArg
}
