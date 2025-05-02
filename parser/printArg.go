package parser

import "fmt"

func (n *Nil) Fmt() string {
	return "NIL"
}

func (a *ArgError) Fmt() string {
	return fmt.Sprintf("ARG_ERR %v", a.Value)
}

func (r *ArgReg) Fmt() string {
	return r.Value.Fmt()
}

func (e *ArgExpr) Fmt() string {
	return e.Value.Fmt()
}
