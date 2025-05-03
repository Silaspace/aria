package language

import "fmt"

func (n *Nil) Fmt() string {
	return "nil"
}

func (e *Error) Fmt() string {
	return fmt.Sprintf("error (%v)", e.Value)
}

func (i *Ident) Fmt() string {
	return fmt.Sprintf("ident (%v)", i.Value)
}

func (r *Reg) Fmt() string {
	return fmt.Sprintf("reg (%v)", r.Value)
}

func (r *RegPair) Fmt() string {
	return fmt.Sprintf("reg (%v : %v)", r.Value+1, r.Value)
}

func (r *RegPointer) Fmt() string {
	return fmt.Sprintf("reg (%v)", r.Value)
}

func (r *RegPointerPostInc) Fmt() string {
	return fmt.Sprintf("reg (%v+)", r.Value)
}

func (r *RegPointerPreDec) Fmt() string {
	return fmt.Sprintf("reg (-%v)", r.Value)
}

func (r *RegPointerDisp) Fmt() string {
	return fmt.Sprintf("reg (%v+%v)", r.Value, r.Disp)
}

func (i *Int) Fmt() string {
	return fmt.Sprintf("int (%v)", i.Value)
}

func (l *List) Fmt() string {
	return fmt.Sprintf("list (%+v)", l.Value)
}

func (a *Assignment) Fmt() string {
	return fmt.Sprintf("assignment (%v = %v)", a.Symbol, a.Value)
}
