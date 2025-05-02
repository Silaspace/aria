package parser

import "fmt"

func (r *Register) Fmt() string {
	return fmt.Sprintf("r%v", r.Value)
}

func (r *RegPair) Fmt() string {
	return fmt.Sprintf("r%v+1:r%v", r.Value, r.Value)
}

func (p *PointerReg) Fmt() string {
	regs := map[RegName]string{X: "x", Y: "y", Z: "z"}
	switch p.Op {
	case None:
		return fmt.Sprintf("%v", regs[p.Value])
	case PostInc:
		return fmt.Sprintf("%v+", regs[p.Value])
	case PreDec:
		return fmt.Sprintf("-%v", regs[p.Value])
	case Disp:
		return fmt.Sprintf("%v+%v", regs[p.Value], p.Disp)
	default:
		return fmt.Sprintf("Unknown operation on pointer register %v", regs[p.Value])
	}
}

func (e *RegErr) Fmt() string {
	return fmt.Sprintf("REG_ERR %v", e.Value)
}
