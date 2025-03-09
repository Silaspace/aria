package language

import (
	"errors"
)

/* -------- Miscellaneous -------- */

/*
Name         A
Description  something
Encoding     0000 0000 1111 1000
*/
func A_5(base uint64, op uint64) (uint64, error) {
	if op > 31 {
		return 0, errors.New("address larger than 5 bits")
	}
	return base | ((op << 3) & 0x00F8), nil
}

/*
Name         A
Description  something
Encoding     0000 0110 0000 1111
*/
func A_6(base uint64, op uint64) (uint64, error) {
	if op > 63 {
		return 0, errors.New("address larger than 6 bits")
	}
	return base | ((op << 5) & 0x0600) | (op & 0x000F), nil
}

/*
Name         b
Description  bit in register
Encoding     0000 0000 0000 0111
*/
func b(base uint64, op uint64) (uint64, error) {
	if op > 7 {
		return 0, errors.New("bit greater than 7")
	}
	return base | (op & 0x0007), nil
}

/*
Name         s
Description  bit in status register
Encoding     0000 0000 0111 0000
*/
func s(base uint64, op uint64) (uint64, error) {
	if op > 7 {
		return 0, errors.New("bit greater than 7")
	}
	return base | ((op << 4) & 0x0030), nil
}

/* -------- Registers -------- */

/*
Name         Rd
Description  desitination register
Encoding     0000 0001 1111 0000
*/
func Rd(base uint64, op uint64) (uint64, error) {
	if op > 31 {
		return 0, errors.New("register specified does not exist")
	}
	return base | ((op << 4) & 0x01F0), nil
}

/*
Name         R
Description  desitination + source register
Encoding     0000 0011 1111 1111
*/
func R(base uint64, op uint64) (uint64, error) {
	if op > 31 {
		return 0, errors.New("register specified does not exist")
	}
	return base | ((op << 4) & 0x01F0) | (op & 0x00F) | ((op << 5) & 0x0200), nil
}

/*
Name         Rd_high
Description  desitination register (r16 to r31)
Encoding     0000 0000 1111 0000
*/
func Rd_high(base uint64, op uint64) (uint64, error) {
	if op > 31 {
		return 0, errors.New("register specified does not exist")
	}
	if op < 16 {
		return 0, errors.New("instructions only operate on the high registers")
	}
	return base | ((op << 4) & 0x00F0), nil
}

/*
Name         R_long
Description  encode register in 32 bit instructions
Encoding     0000 0001 1111 0000 0000 0000 0000 0000
*/
func R_long(base uint64, op uint64) (uint64, error) {
	if op > 31 {
		return 0, errors.New("register specified does not exist")
	}
	return base | ((op << 20) & 0x01F00000), nil
}

/*
Name         Rr
Description  return register
Encoding     0000 0010 0000 1111
*/
func Rr(base uint64, op uint64) (uint64, error) {
	if op > 31 {
		return 0, errors.New("register specified does not exist")
	}
	return base | (op & 0x00F) | ((op << 5) & 0x0200), nil
}

/* -------- Constants -------- */

/*
Name         k_22
Description  22 bit constant for long jmp instructions
Encoding     0000 0001 1111 0001 1111 1111 1111 1111
*/
func k_22(base uint64, op uint64) (uint64, error) {
	if op > 4194304 {
		return 0, errors.New("k larger than 22 bits")
	}
	return base | (op << 25 & 0x01F00000) | (op & 0x0001FFFF), nil
}

/*
Name         k_16
Description  16 bit constant for long instructions
Encoding     1111 1111 1111 1111
*/
func k_16(base uint64, op uint64) (uint64, error) {
	if op > 65535 {
		return 0, errors.New("k larger than 16 bits")
	}
	return base | (op & 0xFFFF), nil
}

/*
Name         k_12
Description  12 bit constant for relative jump
Encoding     0000 1111 1111 1111
*/
func k_12(base uint64, op uint64) (uint64, error) {
	if int16(op) > 4096 {
		return 0, errors.New("k larger than 12 bits")
	}
	return base | (op & 0x0FFF), nil
}

/*
Name         k_8
Description  8 bit constant for immediate
Encoding     0000 1111 0000 1111
*/
func k_8(base uint64, op uint64) (uint64, error) {
	if int16(op) > 255 {
		return 0, errors.New("k larger than 8 bits")
	}
	return base | (op & 0x000F) | ((op << 4) & 0x0F00), nil
}

/*
Name         k_8_compliment
Description  8 bit constant for immediate, complimented
Encoding     0000 1111 0000 1111
*/
func k_8_compliment(base uint64, op uint64) (uint64, error) {
	if int16(op) > 255 {
		return 0, errors.New("k larger than 8 bits")
	}
	comp := op ^ 0xFFFF
	return base | (comp & 0x000F) | ((comp << 4) & 0x0F00), nil
}

/*
Name         k_6
Description  6 bit constant for relative jump
Encoding     0000 0011 1111 1000
*/
func k_6(base uint64, op uint64) (uint64, error) {
	if int16(op) > 64 {
		return 0, errors.New("k larger than 6 bits")
	}
	return base | (op & 0x03F8), nil
}

/*












Addressing Modes

- Single Reg		Rd
- Two Reg			Rd, Rr
- I/O				Rr/Rd, A
- Data Direct		Rr/Rd, Address		[long]
- Data Indirect							[access X/Y/Z Reg]
- Data Ind w/ Pre
- Data Ind w/ Post
- Data Ind w/ Disp	Rr/Rd, q
-



Rd: Destination (and source) register in the Register File
Rr: Source register in the Register File
R: Result after instruction is executed
K: Constant data
k: Constant address
b: Bit position (0..7) in the Register File or I/O Register
s: Bit position (0..7)in the Status Register
X,Y,Z: Indirect Address Register (X=R27:R26, Y=R29:R28, and Z=R31:R30 or X=RAMPX:R27:R26, Y=RAMPY:R29:R28, and Z=RAMPZ:R31:R30 if the memory is larger than 64 KB)
A: I/O memory address
q: Displacement for direct addressing
UU Unsigned × Unsigned operands
SS Signed × Signed operands
SU Signed × Unsigned operands

*/
