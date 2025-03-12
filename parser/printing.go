package parser

import "fmt"

func (e *EOF) Fmt() string {
	return fmt.Sprintf("%v : EOF\n", e.Line)
}

func (e *Error) Fmt() string {
	return fmt.Sprintf("%v : ERROR '%v'\n", e.Line, e.Value)
}

func (c *Comment) Fmt() string {
	return fmt.Sprintf("%v : COM   '%v'\n", c.Line, c.Value)
}

func (l *Label) Fmt() string {
	return fmt.Sprintf("%v : LABEL '%v'\n", l.Line, l.Value)
}

func (d *Directive) Fmt() string {
	dirval := d.Value.Fmt()
	return fmt.Sprintf("%v : DIR '%v'\n", d.Line, dirval)
}

func (i *Instruction) Fmt() string {
	argstr1 := i.Op1.Fmt()
	argstr2 := i.Op2.Fmt()
	return fmt.Sprintf("%v : INSTR '%v' '%v', '%v'\n", i.Line, i.Mnemonic, argstr1, argstr2)
}

func (n *Nil) Fmt() string {
	return "NIL"
}

func (a *ArgError) Fmt() string {
	return "ARG_ERR"
}

func (r *ArgReg) Fmt() string {
	return fmt.Sprintf("r%v", r.Value)
}

func (e *ArgExpr) Fmt() string {
	return e.Value.Fmt()
}

func (e *ErrorExpr) Fmt() string {
	return "EXPR_ERR"
}

func (i *Ident) Fmt() string {
	return i.Value
}

func (l *Literal) Fmt() string {
	return fmt.Sprintf("LIT %v", l.Value)
}

func (b *BinopExpr) Fmt() string {
	e1str := b.E1.Fmt()
	e2str := b.E2.Fmt()
	return fmt.Sprintf("(%v OP %v)", e1str, e2str)
}

func (m *MonopExpr) Fmt() string {
	estr := m.E1.Fmt()
	return fmt.Sprintf("(OP %v)", estr)
}

func (e *ErrorDirVal) Fmt() string {
	return "DIR_ERR"
}

func (n *NilDirVal) Fmt() string {
	return "NIL"
}

func (i *IdentDirVal) Fmt() string {
	return fmt.Sprintf("IDENT %v", i.Value)
}

func (i *ImmDirVal) Fmt() string {
	return fmt.Sprintf("IMM %v", i.Value)
}

func (e *ExprDirVal) Fmt() string {
	estr := e.Value.Fmt()
	return fmt.Sprintf("EXPR %v", estr)
}

func (e *ExprListDirVal) Fmt() string {
	output := []string{}

	for _, expr := range e.Value {
		output = append(output, expr.Fmt())
	}

	return fmt.Sprintf("%v", output)
}

func (e *AssignDirVal) Fmt() string {
	estr := e.Value.Fmt()
	return fmt.Sprintf("%v = %v", e.Symbol, estr)
}
