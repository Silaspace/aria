package parser

import "github.com/silaspace/aria/language"

type ExprType int

type Expr interface {
	Type() ExprType
	Fmt() string
}

const (
	ExprErr   ExprType = 0
	ExprLit   ExprType = 1
	ExprIdent ExprType = 2
	ExprMonop ExprType = 3
	ExprBinop ExprType = 4
	ExprFunc  ExprType = 5
)

type ErrorExpr struct {
	Value string
}

type Ident struct {
	Value string
}

type Literal struct {
	Base  int
	Value string
}

type BinopExpr struct {
	E1     Expr
	E2     Expr
	Symbol string
	Op     language.Operator
}

type MonopExpr struct {
	E1     Expr
	Symbol string
	Op     language.Operator
}

type FuncExpr struct {
	E1     Expr
	Symbol string
	Func   language.Function
}

func (e *ErrorExpr) Type() ExprType {
	return ExprErr
}

func (l *Ident) Type() ExprType {
	return ExprIdent
}

func (l *Literal) Type() ExprType {
	return ExprLit
}

func (b *BinopExpr) Type() ExprType {
	return ExprBinop
}

func (m *MonopExpr) Type() ExprType {
	return ExprMonop
}

func (f *FuncExpr) Type() ExprType {
	return ExprFunc
}
