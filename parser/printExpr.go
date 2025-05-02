package parser

import "fmt"

func (e *ErrorExpr) Fmt() string {
	return fmt.Sprintf("EXPR_ERR %v", e.Value)
}

func (i *Ident) Fmt() string {
	return i.Value
}

func (l *Literal) Fmt() string {
	return fmt.Sprintf("%v", l.Value)
}

func (b *BinopExpr) Fmt() string {
	e1str := b.E1.Fmt()
	e2str := b.E2.Fmt()
	return fmt.Sprintf("(%v %v %v)", e1str, b.Symbol, e2str)
}

func (m *MonopExpr) Fmt() string {
	estr := m.E1.Fmt()
	return fmt.Sprintf("(%v %v)", m.Symbol, estr)
}

func (f *FuncExpr) Fmt() string {
	estr := f.E1.Fmt()
	return fmt.Sprintf("%v (%v)", f.Symbol, estr)
}
