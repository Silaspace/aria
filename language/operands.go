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
	-- LD --
	100b 0000 000 ppff

	d - Rd
	p - Z (1) / Y (2) / X (3)
	f - None (0) / PostInc (1) / PreDec (2)
	b = (Y | Z) & None
*/

/*
Name         ld_pointer
Description  encodes X, Y, or Z in unchanged, postinc or predec form for LD
Encoding

	()   100b 0000 0000 ppff

	(X)  1001 0000 0000 1100
	(X+) 1001 0000 0000 1101
	(-X) 1001 0000 0000 1110

	(Y)  1000 0000 0000 1000
	(Y+) 1001 0000 0000 1001
	(-Y) 1001 0000 0000 1010

	(Z)  1000 0000 0000 0000
	(Z+) 1001 0000 0000 0001
	(-Z) 1001 0000 0000 0010
*/
func ld_pointer(base uint64, rp Value) (uint64, error) {
	switch rp := rp.(type) {
	case *RegPointer:
		reg := base >> 4

		if rp.Op == PostInc || rp.Op == PreDec {
			switch rp.Value {
			case X:
				if reg == 26 || reg == 27 {
					return 0, errors.New("ld (r26|r27) (x+|-x) is undefined")
				}
			case Y:
				if reg == 28 || reg == 29 {
					return 0, errors.New("ld (r28|r29) (y+|-y) is undefined")
				}
			case Z:
				if reg == 30 || reg == 31 {
					return 0, errors.New("ld (r30|r31) (z+|-z) is undefined")
				}
			}
		}

		/*
			p = X (3)
			  | Y (2)
			  | Z (0)
		*/
		if rp.Value == X {
			base = base | 0x000C
		}

		if rp.Value == Y {
			base = base | 0x0008
		}

		/*
			b = not((Y or Z) and None)
			  = X or (not None)
		*/
		if (rp.Value == X) && (rp.Op != None) {
			base = base | 0x1000
		}

		/*
			f = None    (0)
			  | PostInc (1)
			  | PreDec  (2)
		*/
		return base | uint64(rp.Op), nil

	case *Error:
		return 0, errors.New(rp.Value)

	default:
		return 0, fmt.Errorf("expected reg pointer, got %+v", rp.Fmt())
	}
}

/*
Name         ldd_pointer
Description  encodes Y or Z in displacement form for LDD
Encoding

	(Y)   10q0 qq00 0000 1qqq
	(Z)   10q0 qq0d dddd 0qqq
*/
func ldd_pointer(base uint64, rp Value) (uint64, error) {
	switch rp := rp.(type) {
	case *RegPointer:
		if rp.Disp > 63 {
			return 0, errors.New("displacement larger than 6 bits")
		}

		if rp.Value == X {
			return 0, errors.New("displacement from X not supported")
		}

		if rp.Value == Y {
			base = base | 0x0008
		}

		return base |
			((rp.Disp << 8) & 0x2000) |
			((rp.Disp << 7) & 0x0C00) |
			(rp.Disp & 0x0007), nil

	case *Error:
		return 0, errors.New(rp.Value)

	default:
		return 0, fmt.Errorf("expected reg pointer, got %+v", rp.Fmt())
	}
}

/*
Name         st_pointer
Description  encodes X, Y, or Z in unchanged, postinc or predec form for ST
Encoding

	(-)  100b 0010 0000 ppff

	(X)  1001 001r rrrr 1100
	(X+) 1001 001r rrrr 1101
	(-X) 1001 001r rrrr 1110

	(Y)  1000 001r rrrr 1000
	(Y+) 1000 001r rrrr 1000
	(-Y) 1001 001r rrrr 1010

	(Z)  1000 001r rrrr 0000
	(Z+) 1001 001r rrrr 0001
	(-Z) 1001 001r rrrr 0010
*/
func st_pointer(base uint64, rp Value) (uint64, error) {
	switch rp := rp.(type) {
	case *RegPointer:
		/*
			p = X (3)
			  | Y (2)
			  | Z (0)
		*/
		if rp.Value == X {
			base = base | 0x000C
		}

		if rp.Value == Y {
			base = base | 0x0008
		}

		/*
			b = not((Y or Z) and None)
			  = X or (not None)
		*/
		if (rp.Value == X) && (rp.Op != None) {
			base = base | 0x1000
		}

		/*
			f = None    (0)
			  | PostInc (1)
			  | PreDec  (2)
		*/
		return base | uint64(rp.Op), nil

	case *Error:
		return 0, errors.New(rp.Value)

	default:
		return 0, fmt.Errorf("expected reg pointer, got %+v", rp.Fmt())
	}
}

/*
Name         std_pointer
Description  encodes Y or Z in displacement form for STD
Encoding

	(Y) 10q0 qq10 0000 1qqq
	(Z) 10q0 qq10 0000 0qqq

	12 82 -> 1000 0010 0001 0010
*/
func std_pointer(base uint64, rp Value) (uint64, error) {
	switch rp := rp.(type) {
	case *RegPointer:
		if rp.Disp > 63 {
			return 0, errors.New("displacement larger than 6 bits")
		}

		if rp.Value == X {
			return 0, errors.New("displacement from X not supported")
		}

		if rp.Value == Y {
			base = base | 0x0008
		}

		return base |
			((rp.Disp << 8) & 0x2000) |
			((rp.Disp << 7) & 0x0C00) |
			(rp.Disp & 0x0007), nil

	case *Error:
		return 0, errors.New(rp.Value)

	default:
		return 0, fmt.Errorf("expected reg pointer, got %+v", rp.Fmt())
	}
}

/*
Name         Rd
Description  desitination register (with undefined check for ST)
Encoding     0000 000d dddd 0000
*/
func Rd_st(base uint64, op Value) (uint64, error) {
	switch op := op.(type) {
	case *Reg:

		if op.Value > 31 {
			return 0, errors.New("register specified does not exist")
		}

		// Check undefinedness
		rpOp := base & 0x0003
		rpValue := (base >> 2) & 0x0003

		if rpOp != 0 {
			switch rpValue {
			case 3: // X
				if op.Value == 26 || op.Value == 27 {
					return 0, errors.New("st (x+|-x) (r26|r27) is undefined")
				}
			case 2: // Y
				if op.Value == 28 || op.Value == 29 {
					return 0, errors.New("st (y+|-y) (r2|r29) is undefined")
				}
			case 0: // Z
				if op.Value == 30 || op.Value == 31 {
					return 0, errors.New("st (z+|-z) (r30|r31) is undefined")
				}
			}
		}

		return base | ((op.Value << 4) & 0x01F0), nil

	case *Error:
		return 0, errors.New(op.Value)

	default:
		return 0, fmt.Errorf("expected reg, got %v", op.Fmt())
	}
}
