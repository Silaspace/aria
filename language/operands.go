package language

import (
	"errors"
	"fmt"
)

/* ------- Named Registers ------- */

type Register struct{}

const (
	PC Mnemonic = "pc"
	X  Mnemonic = "x"
	Y  Mnemonic = "y"
	Z  Mnemonic = "z"
)

var Registers = map[Mnemonic]Register{
	PC: {},
	X:  {},
	Y:  {},
	Z:  {},
}

/* -------- Miscellaneous -------- */

/*
Name         A
Description  something
Encoding     0000 0000 1111 1000
*/
func A_5(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if op.Value > 31 {
			return 0, errors.New("address larger than 5 bits")
		}
		return base | ((op.Value << 3) & 0x00F8), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         A
Description  something
Encoding     0000 0110 0000 1111
*/
func A_6(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if op.Value > 63 {
			return 0, errors.New("address larger than 6 bits")
		}
		return base | ((op.Value << 5) & 0x0600) | (op.Value & 0x000F), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         b
Description  bit in register
Encoding     0000 0000 0000 0111
*/
func b(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if op.Value > 7 {
			return 0, errors.New("bit greater than 7")
		}
		return base | (op.Value & 0x0007), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         s
Description  bit in status register
Encoding     0000 0000 0111 0000
*/
func s(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if op.Value > 7 {
			return 0, errors.New("bit greater than 7")
		}
		return base | ((op.Value << 4) & 0x0030), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/* -------- Registers -------- */

/*
Name         Rd
Description  desitination register
Encoding     0000 0001 1111 0000
*/
func Rd(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Reg:
		if op.Value > 31 {
			return 0, errors.New("register specified does not exist")
		}
		return base | ((op.Value << 4) & 0x01F0), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected reg, got %v", op.Fmt())
	}
}

/*
Name         R
Description  desitination + source register
Encoding     0000 0011 1111 1111
*/
func R(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Reg:
		if op.Value > 31 {
			return 0, errors.New("register specified does not exist")
		}
		return base | ((op.Value << 4) & 0x01F0) | (op.Value & 0x00F) | ((op.Value << 5) & 0x0200), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected reg, got %v", op.Fmt())
	}
}

/*
Name         Rd_high
Description  desitination register (r16 to r31)
Encoding     0000 0000 1111 0000
*/
func Rd_high(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Reg:
		if op.Value > 31 {
			return 0, errors.New("register specified does not exist")
		}
		if op.Value < 16 {
			return 0, errors.New("instructions only operate on the high registers")
		}
		return base | ((op.Value << 4) & 0x00F0), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected reg, got %+v", op.Fmt())
	}
}

/*
Name         R_long
Description  encode register in 32 bit instructions
Encoding     0000 0001 1111 0000 0000 0000 0000 0000
*/
func R_long(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Reg:
		if op.Value > 31 {
			return 0, errors.New("register specified does not exist")
		}
		return base | ((op.Value << 20) & 0x01F00000), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected reg, got %+v", op.Fmt())
	}
}

/*
Name         Rr
Description  return register
Encoding     0000 0010 0000 1111
*/
func Rr(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Reg:
		if op.Value > 31 {
			return 0, errors.New("register specified does not exist")
		}
		return base | (op.Value & 0x00F) | ((op.Value << 5) & 0x0200), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected reg, got %v", op.Fmt())
	}
}

/* -------- Constants -------- */

/*
Name         k_22
Description  22 bit constant for long jmp instructions
Encoding     0000 0001 1111 0001 1111 1111 1111 1111
*/
func k_22(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if op.Value > 4194304 {
			return 0, errors.New("k larger than 22 bits")
		}
		return base | (op.Value << 25 & 0x01F00000) | (op.Value & 0x0001FFFF), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         k_16
Description  16 bit constant for long instructions
Encoding     1111 1111 1111 1111
*/
func k_16(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if op.Value > 65535 {
			return 0, errors.New("k larger than 16 bits")
		}
		return base | (op.Value & 0xFFFF), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         k_12
Description  12 bit constant for relative jump
Encoding     0000 1111 1111 1111
*/
func k_12(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if int16(op.Value) > 4096 {
			return 0, errors.New("k larger than 12 bits")
		}
		return base | (op.Value & 0x0FFF), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         k_8
Description  8 bit constant for immediate
Encoding     0000 1111 0000 1111
*/
func k_8(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if int16(op.Value) > 255 {
			return 0, errors.New("k larger than 8 bits")
		}
		return base | (op.Value & 0x000F) | ((op.Value << 4) & 0x0F00), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         k_8_compliment
Description  8 bit constant for immediate, complimented
Encoding     0000 1111 0000 1111
*/
func k_8_compliment(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if int16(op.Value) > 255 {
			return 0, errors.New("k larger than 8 bits")
		}
		comp := op.Value ^ 0xFFFF
		return base | (comp & 0x000F) | ((comp << 4) & 0x0F00), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         k_6
Description  6 bit constant for relative jump
Encoding     0000 0011 1111 1000
*/
func k_6(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if int16(op.Value) > 64 {
			return 0, errors.New("k larger than 6 bits")
		}
		return base | (op.Value & 0x03F8), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}
