package language

import "errors"

func (n *Nil) Augment(v Value) error {
	return nil
}

func (e *Error) Augment(v Value) error {
	return nil
}

func (i *Ident) Augment(v Value) error {
	return nil
}

func (r *Reg) Augment(v Value) error {
	return nil
}

func (r *RegPair) Augment(v Value) error {
	return nil
}

func (r *RegPointer) Augment(v Value) error {
	return nil
}

func (r *RegPointerPostInc) Augment(v Value) error {
	switch v := v.(type) {
	case *Reg:
		r.Reg = *v
		return nil

	default:
		return errors.New("expected reg argument in augmentation")
	}
}

func (r *RegPointerPreDec) Augment(v Value) error {
	switch v := v.(type) {
	case *Reg:
		r.Reg = *v
		return nil

	default:
		return errors.New("expected reg argument in augmentation")
	}
}

func (r *RegPointerDisp) Augment(v Value) error {
	return nil
}

func (i *Int) Augment(v Value) error {
	return nil
}

func (l *List) Augment(v Value) error {
	return nil
}

func (a *Assignment) Augment(v Value) error {
	return nil
}
