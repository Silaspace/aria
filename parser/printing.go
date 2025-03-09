package parser

import "fmt"

func (e *EOF) Fmt() string {
	return "EOF\n"
}

func (e *Error) Fmt() string {
	return fmt.Sprintf("ERROR '%v'\n", e.Value)
}

func (c *Comment) Fmt() string {
	return fmt.Sprintf("COM   '%v'\n", c.Value)
}

func (l *Label) Fmt() string {
	return fmt.Sprintf("LABEL '%v'\n", l.Value)
}

func (d *Directive) Fmt() string {
	return "DIR\n"
}

func (i *Instruction) Fmt() string {
	argstr1 := i.Op1.Fmt()
	argstr2 := i.Op2.Fmt()
	return fmt.Sprintf("INSTR '%v' '%v', '%v'\n", i.Mnemonic, argstr1, argstr2)
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
