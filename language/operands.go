package language

import (
	"errors"
	"fmt"
)

/* ------- Named Registers ------- */

type Register struct{}

type PointerKey struct {
	instr Mnemonic
	reg   Mnemonic
	op    PointerOp
}

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

var PointerInstructions = map[PointerKey]uint64{
	{LD, X, None}:    0x900C, // 1001 0000 0000 1100
	{LD, X, PostInc}: 0x900D, // 1001 0000 0000 1101
	{LD, X, PreDec}:  0x900E, // 1001 0000 0000 1110
	{LD, Y, None}:    0x800D, // 1000 0000 0000 1000
	{LD, Y, PostInc}: 0x9009, // 1001 0000 0000 1001
	{LD, Y, PreDec}:  0x900A, // 1001 0000 0000 1010
	{LD, Z, None}:    0x8000, // 1000 0000 0000 0000
	{LD, Z, PostInc}: 0x9001, // 1001 0000 0000 0001
	{LD, Z, PreDec}:  0x9002, // 1001 0000 0000 0010
	{LDD, Y, Disp}:   0x8008, // 10q0 qq0d dddd 1qqq
	{LDD, Z, Disp}:   0x8000, // 10q0 qq0d dddd 0qqq
}

/* -------- Miscellaneous -------- */

/*
Name         A
Description  5-bit memory address
Encoding     0000 0000 aaaa a000
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
Description  6-bit memory address
Encoding     0000 0aa0 0000 aaaa
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
Encoding     0000 0000 0000 0bbb
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
Encoding     0000 0000 0sss 0000
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
Encoding     0000 000d dddd 0000
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
Encoding     0000 00rd dddd rrrr
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
Encoding     0000 0000 dddd 0000
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
Encoding     0000 000r rrrr 0000 0000 0000 0000 0000
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
Encoding     0000 00r0 0000 rrrr
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

/*
Name         Rd+1:Rd
Description  upper register pairs - d ∈ {24,26,28,30}
Encoding     0000 0000 00dd 0000
*/
func R_pair(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *RegPair:
		switch op.Value {
		case 24, 26, 28, 30:
			return base | ((op.Value << 3) & 0x0030), nil

		default:
			return 0, errors.New("register pair specified does not exist")
		}

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected reg pair, got %v", op.Fmt())
	}
}

/* -------- Constants -------- */

/*
Name         k_22
Description  22 bit constant for long jmp instructions
Encoding     0000 000k kkkk 000k kkkk kkkk kkkk kkkk
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
Encoding     kkkk kkkk kkkk kkkk
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
Encoding     0000 kkkk kkkk kkkk
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
Encoding     0000 kkkk 0000 kkkk
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
Encoding     0000 kkkk 0000 kkkk
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
Encoding     0000 00kk kkkk k000
*/
func k_6(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if int16(op.Value) > 63 {
			return 0, errors.New("k larger than 6 bits")
		}
		return base | (op.Value & 0x03F8), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/*
Name         k_6_ii
Description  6 bit constant for add immediate word (ADIW)
Encoding     0000 0000 kk00 kkkk
*/
func k_6_ii(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Int:
		if int16(op.Value) > 63 {
			return 0, errors.New("k larger than 6 bits")
		}
		return base | (op.Value & 0x00CF), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected int, got %+v", op.Fmt())
	}
}

/* -------- Pointer Register Operands -------- */

/*
Name         Pointer_ld
Description  encodes X, Y, or Z in unchanged, postinc or predec form
Encoding

	(i)    1001 0000 0000 1100
	(ii)   1001 0000 0000 1101
	(iii)  1001 0000 0000 1110

	(iv)   1000 0000 0000 1000
	(v)    1001 0000 0000 1001
	(vi)   1001 0000 0000 1010

	(vi)   1000 0000 0000 0000
	(viii) 1001 0000 0000 0001
	(ix)   1001 0000 0000 0010
*/
func Pointer_ld(base uint64, rp Value) (uint64, error) {
	switch rp := rp.(type) {
	case *RegPointer:
		reg := base >> 4

		if rp.Op == PostInc || rp.Op == PreDec {
			switch rp.Value {
			case X:
				if reg == 26 || reg == 27 {
					return 0, errors.New("ld r26 x+ and ld r27 x+ are undefined")
				}
			case Y:
				if reg == 28 || reg == 29 {
					return 0, errors.New("ld r28 y+ and ld r29 y+ are undefined")
				}
			case Z:
				if reg == 30 || reg == 31 {
					return 0, errors.New("ld r30 z+ and ld r31 z+ are undefined")
				}
			}
		}

		instr, exists := PointerInstructions[PointerKey{LD, rp.Value, rp.Op}]

		if !exists {
			return 0, fmt.Errorf("(%v, %v, %v) is an undefined operation", LD, rp.Value, rp.Op)
		}

		return instr, nil

	case *Error:
		return 0, errors.New(rp.Value)

	default:
		return 0, fmt.Errorf("expected reg pointer, got %+v", rp.Fmt())
	}
}

/*
Name         Disp_ld
Description  encodes Y or Z in displacement form
Encoding

	(i)    10q0 qq0d dddd 1qqq
	(ii)   10q0 qq0d dddd 0qqq
*/
func Disp_ld(base uint64, rp Value) (uint64, error) {
	switch rp := rp.(type) {
	case *RegPointer:
		instr, exists := PointerInstructions[PointerKey{LDD, rp.Value, rp.Op}]

		if !exists {
			return 0, fmt.Errorf("(%v, %v, %v) is an undefined operation", LD, rp.Value, rp.Op)
		}

		return instr, nil

	case *Error:
		return 0, errors.New(rp.Value)

	default:
		return 0, fmt.Errorf("expected reg pointer, got %+v", rp.Fmt())
	}
}
