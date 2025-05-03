package parser

import (
	"fmt"
)

func (r *Register) Fmt() string {
	return fmt.Sprintf("r%v", r.Value)
}

func (r *RegPair) Fmt() string {
	return fmt.Sprintf("r%v+1:r%v", r.Value, r.Value)
}

func (p *RegPointer) Fmt() string {
	return fmt.Sprintf("%v", p.Value)
}

func (p *RegPointerPostInc) Fmt() string {
	return fmt.Sprintf("%v+", p.Value)
}

func (p *RegPointerPreDec) Fmt() string {
	return fmt.Sprintf("-%v", p.Value)
}

func (p *RegPointerDisp) Fmt() string {
	return fmt.Sprintf("%v+%v", p.Value, p.Disp)
}

func (e *RegErr) Fmt() string {
	return fmt.Sprintf("REG_ERR %v", e.Value)
}
