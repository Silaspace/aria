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
