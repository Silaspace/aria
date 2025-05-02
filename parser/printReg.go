package parser

import (
	"fmt"

	"github.com/silaspace/aria/language"
)

func (r *Register) Fmt() string {
	return fmt.Sprintf("r%v", r.Value)
}

func (r *RegPair) Fmt() string {
	return fmt.Sprintf("r%v+1:r%v", r.Value, r.Value)
}

func (p *PointerReg) Fmt() string {
	switch p.Op {
	case language.None:
		return fmt.Sprintf("%v", p.Value)
	case language.PostInc:
		return fmt.Sprintf("%v+", p.Value)
	case language.PreDec:
		return fmt.Sprintf("-%v", p.Value)
	case language.Disp:
		return fmt.Sprintf("%v+%v", p.Value, p.Disp)
	default:
		return fmt.Sprintf("Unknown operation on pointer register %v", p.Value)
	}
}

func (e *RegErr) Fmt() string {
	return fmt.Sprintf("REG_ERR %v", e.Value)
}
