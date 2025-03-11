package parser

type DirValType int

type DirVal interface {
	Type() DirValType
	Fmt() string
}

const (
	DirValErr      DirValType = 0
	DirValNil      DirValType = 1
	DirValIdent    DirValType = 2
	DirValImm      DirValType = 3
	DirValExpr     DirValType = 4
	DirValExprList DirValType = 5
	DirValAssign   DirValType = 6
)

type ErrorDirVal struct {
	Value string
}

type NilDirVal struct{}

type IdentDirVal struct {
	Value string
}

type ImmDirVal struct {
	Value string
}

type ExprDirVal struct {
	Value Expr
}

type ExprListDirVal struct {
	Value []Expr
}

type AssignDirVal struct {
	Symbol string
	Value  Expr
}

func (e *ErrorDirVal) Type() DirValType {
	return DirValErr
}

func (e *NilDirVal) Type() DirValType {
	return DirValNil
}

func (e *IdentDirVal) Type() DirValType {
	return DirValIdent
}

func (e *ImmDirVal) Type() DirValType {
	return DirValImm
}

func (e *ExprDirVal) Type() DirValType {
	return DirValExpr
}

func (e *ExprListDirVal) Type() DirValType {
	return DirValExprList
}

func (e *AssignDirVal) Type() DirValType {
	return DirValAssign
}
