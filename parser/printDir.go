package parser

import "fmt"

func (e *ErrorDirVal) Fmt() string {
	return fmt.Sprintf("DIR_ERR %v", e.Value)
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
