package language

import (
	"encoding/binary"
	"fmt"
)

type Flag int
type OpFunc func(uint64, uint64) (uint64, error)

type Instruction struct {
	Mnemonic Mnemonic
	Base     uint64
	Op1      OpFunc
	Op2      OpFunc
	Flags    Flag
}

const (
	LONG     Flag = 1
	RELATIVE Flag = 1 << 1
)

const (
	ADC   Mnemonic = "adc"
	ADD   Mnemonic = "add"
	AND   Mnemonic = "and"
	ANDI  Mnemonic = "andi"
	ADIW  Mnemonic = "andw" /* AVRe */
	ASR   Mnemonic = "asr"
	BCLR  Mnemonic = "bclr"
	BLD   Mnemonic = "bld"
	BRBC  Mnemonic = "brbc"
	BRBS  Mnemonic = "brbs"
	BRCC  Mnemonic = "brcc"
	BRCS  Mnemonic = "brcs"
	BREAK Mnemonic = "break" /* AVRe */
	BREQ  Mnemonic = "breq"
	BRGE  Mnemonic = "brge"
	BRHC  Mnemonic = "brhc"
	BRHS  Mnemonic = "brhs"
	BRID  Mnemonic = "brid"
	BIRE  Mnemonic = "brie"
	BRLO  Mnemonic = "brlo"
	BRLT  Mnemonic = "brlt"
	BRMI  Mnemonic = "brmi"
	BRNE  Mnemonic = "brne"
	BRPL  Mnemonic = "brpl"
	BRSH  Mnemonic = "brsh"
	BRTC  Mnemonic = "brtc"
	BRTS  Mnemonic = "brts"
	BRVC  Mnemonic = "brvc"
	BRVS  Mnemonic = "brvs"
	BSET  Mnemonic = "bset"
	BST   Mnemonic = "bst"
	CALL  Mnemonic = "call" /* AVRe */
	CBI   Mnemonic = "cbi"
	CBR   Mnemonic = "cbr"
	CLC   Mnemonic = "clc"
	CLH   Mnemonic = "clh"
	CLI   Mnemonic = "cli"
	CLN   Mnemonic = "cln"
	CLR   Mnemonic = "clr"
	CLS   Mnemonic = "cls"
	CLT   Mnemonic = "clt"
	CLV   Mnemonic = "clv"
	CLZ   Mnemonic = "clz"
	COM   Mnemonic = "com"
	CP    Mnemonic = "cp"
	CPC   Mnemonic = "cpc"
	CPI   Mnemonic = "cpi"
	CPSE  Mnemonic = "cpse"
	DEC   Mnemonic = "dec"
	EOR   Mnemonic = "eor"
	ICALL Mnemonic = "icall"
	IJMP  Mnemonic = "ijmp"
	IN    Mnemonic = "in"
	INC   Mnemonic = "inc"
	JMP   Mnemonic = "jmp" /* AVRe */
	LD    Mnemonic = "ld"
	LDD   Mnemonic = "ld" /* AVRe */
	LDI   Mnemonic = "ldi"
	LDS   Mnemonic = "lds"
	LPM   Mnemonic = "lpm" /* AVRe */
	LSL   Mnemonic = "lsl"
	LSR   Mnemonic = "lsr"
	MOV   Mnemonic = "mov"
	NEG   Mnemonic = "neg"
	NOP   Mnemonic = "nop"
	OR    Mnemonic = "or"
	ORI   Mnemonic = "ori"
	OUT   Mnemonic = "out"
	POP   Mnemonic = "pop"
	PUSH  Mnemonic = "push"
	RCALL Mnemonic = "rcall"
	RET   Mnemonic = "ret"
	RETI  Mnemonic = "reti"
	RJMP  Mnemonic = "rjmp"
	ROL   Mnemonic = "rol"
	ROR   Mnemonic = "ror"
	SBC   Mnemonic = "sbc"
	SBCI  Mnemonic = "sbci"
	SBI   Mnemonic = "sbi"
	SBIC  Mnemonic = "sbic"
	SBIS  Mnemonic = "sbis"
	SBIW  Mnemonic = "sbiw" /* AVRe */
	SBR   Mnemonic = "sbr"
	SBRC  Mnemonic = "sbrc"
	SBRS  Mnemonic = "sbrs"
	SEC   Mnemonic = "sec"
	SEH   Mnemonic = "seh"
	SEI   Mnemonic = "sei"
	SEN   Mnemonic = "sen"
	SER   Mnemonic = "ser"
	SES   Mnemonic = "ses"
	SET   Mnemonic = "set"
	SEV   Mnemonic = "sev"
	SEZ   Mnemonic = "sez"
	SLEEP Mnemonic = "sleep"
	SPM   Mnemonic = "spm" /* AVRe */
	ST    Mnemonic = "st"
	STD   Mnemonic = "std" /* AVRe */
	STS   Mnemonic = "sts"
	SUB   Mnemonic = "sub"
	SUBI  Mnemonic = "subi"
	SWAP  Mnemonic = "swap"
	TST   Mnemonic = "tst"
	WDR   Mnemonic = "wdr"
)

var AVR = map[Mnemonic]Instruction{
	/*
		Syntax    ADC Rd, Rr
		Encoding  0001 11rd dddd rrrr
	*/
	ADC: {
		Base:  0x1c00,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    ADD Rd, Rr
		Encoding  0000 11rd dddd rrrr
	*/
	ADD: {
		Base:  0x0c00,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    AND Rd, Rr
		Encoding  0010 00rd dddd rrrr
	*/
	AND: {
		Base:  0x2000,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    ANDI Rd, K
		Encoding  0111 KKKK dddd KKKK
	*/
	ANDI: {
		Base:  0x7000,
		Op1:   Rd_high,
		Op2:   k_8,
		Flags: 0,
	},

	/*
		Syntax    ASR Rd
		Encoding  1001 010d dddd 0101
	*/
	ASR: {
		Base:  0x0000,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax   BCLR s
		Encoding 1001 0100 1sss 1000
	*/
	BCLR: {
		Base:  0x9488,
		Op1:   s,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    BLD Rd, b
		Encoding  1111 100d dddd 0bbb
	*/
	BLD: {
		Base:  0xF800,
		Op1:   Rd,
		Op2:   b,
		Flags: 0,
	},

	/*
		Syntax    BRBC s, k
		Encoding  1111 01kk kkkk ksss
	*/
	BRBC: {
		Base:  0xF400,
		Op1:   b,
		Op2:   k_6,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRBS s, k
		Encoding  1111 00kk kkkk ksss
	*/
	BRBS: {
		Base:  0xF000,
		Op1:   b,
		Op2:   k_6,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRCC k
		Encoding  1111 01kk kkkk k000
	*/
	BRCC: {
		Base:  0xF400,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRCS k
		Encoding  1111 00kk kkkk k000
	*/
	BRCS: {
		Base:  0xF000,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BREQ k
		Encoding  1111 00kk kkkk k001
	*/
	BREQ: {
		Base:  0xF001,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRGE k
		Encoding  1111 01kk kkkk k100
	*/
	BRGE: {
		Base:  0xF401,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRHC k
		Encoding  1111 01kk kkkk k101
	*/
	BRHC: {
		Base:  0xF405,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRHS k
		Encoding  1111 00kk kkkk k101
	*/
	BRHS: {
		Base:  0xF005,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRID k
		Encoding  1111 01kk kkkk k111
	*/
	BRID: {
		Base:  0xF407,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRIE k
		Encoding  1111 00kk kkkk k111
	*/
	BIRE: {
		Base:  0xF007,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRLO k
		Encoding  1111 00kk kkkk k000
	*/
	BRLO: {
		Base:  0xF000,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRLT k
		Encoding  1111 00kk kkkk k100
	*/
	BRLT: {
		Base:  0xF004,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRMI k
		Encoding  1111 00kk kkkk k010
	*/
	BRMI: {
		Base:  0xF002,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRNE k
		Encoding  1111 01kk kkkk k001
	*/
	BRNE: {
		Base:  0xF401,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRPL k
		Encoding  1111 01kk kkkk k010
	*/
	BRPL: {
		Base:  0xF402,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRSH k
		Encoding  1111 01kk kkkk k000
	*/
	BRSH: {
		Base:  0xF400,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRTC k
		Encoding  1111 01kk kkkk k110
	*/
	BRTC: {
		Base:  0xF406,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRTS k
		Encoding  1111 00kk kkkk k110
	*/
	BRTS: {
		Base:  0xF006,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRVC k
		Encoding  1111 01kk kkkk k011
	*/
	BRVC: {
		Base:  0xF403,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BRVS k
		Encoding  1111 00kk kkkk k011
	*/
	BRVS: {
		Base:  0xF003,
		Op1:   k_6,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    BSET s
		Encoding  1001 0100 0sss 1000
	*/
	BSET: {
		Base:  0x9408,
		Op1:   s,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    BST Rd, b
		Encoding  1111 101d dddd 0bbb
	*/
	BST: {
		Base:  0xFA00,
		Op1:   Rd,
		Op2:   b,
		Flags: 0,
	},

	/*
		Syntax    CBI A, b
		Encoding  1001 1000 AAAA Abbb
	*/
	CBI: {
		Base:  0x9800,
		Op1:   A_5,
		Op2:   b,
		Flags: 0,
	},

	/*
		Syntax    CBR Rd, k
		Encoding  0111 KKKK dddd KKKK
	*/
	CBR: {
		Base:  0x7000,
		Op1:   Rd,
		Op2:   k_8_compliment,
		Flags: 0,
	},

	/*
		Syntax    CLC
		Encoding  1001 0100 1000 1000
	*/
	CLC: {
		Base:  0x9488,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CLH
		Encoding  1001 0100 1101 1000
	*/
	CLH: {
		Base:  0x94D8,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CLI
		Encoding  1001 0100 1111 1000
	*/
	CLI: {
		Base:  0x94F8,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CLN
		Encoding  1001 0100 1010 1000
	*/
	CLN: {
		Base:  0x94A8,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CLR Rd
		Encoding  0010 01dd dddd dddd
	*/
	CLR: {
		Base:  0x2400,
		Op1:   R,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CLS
		Encoding  1001 0100 1100 1000
	*/
	CLS: {
		Base:  0x94C8,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CLT
		Encoding  1001 0100 1110 1000
	*/
	CLT: {
		Base:  0x94E8,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CLV
		Encoding  1001 0100 1011 1000
	*/
	CLV: {
		Base:  0x94B8,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CLZ
		Encoding  1001 0100 1001 1000
	*/
	CLZ: {
		Base:  0x9498,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    COM Rd
		Encoding  1001 010d dddd 0000
	*/
	COM: {
		Base:  0x9400,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    CP Rd, Rr
		Encoding  0001 01rd dddd rrrr
	*/
	CP: {
		Base:  0x1400,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    CPC Rd, Rr
		Encoding  0000 01rd dddd rrrr
	*/
	CPC: {
		Base:  0x0400,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    CPI Rd, K
		Encoding  0011 KKKK dddd KKKK
	*/
	CPI: {
		Base:  0x3000,
		Op1:   Rd_high,
		Op2:   k_8,
		Flags: 0,
	},

	/*
		Syntax    CPSE Rd, Rr
		Encoding  0001 00rd dddd rrrr
	*/
	CPSE: {
		Base:  0x1000,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    DEC Rd
		Encoding  1001 010d dddd 1010
	*/
	DEC: {
		Base:  0x940A,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    EOR Rd, Rr
		Encoding  0010 01rd dddd rrrr
	*/
	EOR: {
		Base:  0x2400,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    ICALL
		Encoding  1001 0101 0000 1001
	*/
	ICALL: {
		Base:  0x9609,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    IJMP
		Encoding  1001 0100 0000 1001
	*/
	IJMP: {
		Base:  0x9409,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    IN Rd, A
		Encoding  1011 0AAd dddd AAAA
	*/
	IN: {
		Base:  0xB000,
		Op1:   Rd,
		Op2:   A_6,
		Flags: 0,
	},

	/*
		Syntax    INC Rd
		Encoding  1001 010d dddd 0011
	*/
	INC: {
		Base:  0x9403,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*  TODO - ELEVEN INSTRUCTION VARIENTS
		Syntax
		Encoding

	LD: {
		Base:  0x0000,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},*/

	/*
		Syntax    LDI Rd, K
		Encoding  1110 KKKK dddd KKKK
	*/
	LDI: {
		Base:  0x0000,
		Op1:   Rd_high,
		Op2:   k_8,
		Flags: 0,
	},

	/*
		Syntax    LDS Rd,k
		Encoding  1001 000d dddd 0000 kkkk kkkk kkkk kkkk
	*/
	LDS: {
		Base:  0x90000000,
		Op1:   R_long,
		Op2:   k_16,
		Flags: LONG,
	},

	/*
		Syntax    LSL Rd
		Encoding  0000 11dd dddd dddd
	*/
	LSL: {
		Base:  0x0C00,
		Op1:   R,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    LSR Rd
		Encoding  1001 010d dddd 0110
	*/
	LSR: {
		Base:  0x9406,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    MOV Rd, Rr
		Encoding  0010 11rd dddd rrrr
	*/
	MOV: {
		Base:  0x2C00,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    NEG Rd
		Encoding  1001 010d dddd 0001
	*/
	NEG: {
		Base:  0x9401,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    NOP
		Encoding  0000 0000 0000 0000
	*/
	NOP: {
		Base:  0x0000,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    OR Rd, Rr
		Encoding  0010 10rd dddd rrrr
	*/
	OR: {
		Base:  0x4800,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    ORI Rd, K
		Encoding  0110 KKKK dddd KKKK
	*/
	ORI: {
		Base:  0x6000,
		Op1:   Rd_high,
		Op2:   k_8,
		Flags: 0,
	},

	/*
		Syntax    OUT A, Rd
		Encoding  1011 1AAr rrrr AAAA
	*/
	OUT: {
		Base:  0xB800,
		Op1:   A_6,
		Op2:   Rd,
		Flags: 0,
	},

	/*
		Syntax    POP Rd
		Encoding  1001 000d dddd 1111
	*/
	POP: {
		Base:  0x900F,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    PUSH Rr
		Encoding  1001 001d dddd 1111
	*/
	PUSH: {
		Base:  0x9200,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax
		Encoding  1101 kkkk kkkk kkkk
	*/
	RCALL: {
		Base:  0xD000,
		Op1:   k_12,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    RET
		Encoding  1001 0101 0000 1000
	*/
	RET: {
		Base:  0x9508,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    RETI
		Encoding  1001 0101 0001 1000
	*/
	RETI: {
		Base:  0x9518,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    RJMP k
		Encoding  1100 kkkk kkkk kkkk
	*/
	RJMP: {
		Base:  0xc000,
		Op1:   k_12,
		Op2:   nil,
		Flags: RELATIVE,
	},

	/*
		Syntax    ROL Rd
		Encoding  0001 11dd dddd dddd
	*/
	ROL: {
		Base:  0x1C00,
		Op1:   R,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    ROR Rd
		Encoding  1001 010d dddd 0111
	*/
	ROR: {
		Base:  0x9407,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SBC Rd, Rr
		Encoding  0000 10rd dddd rrrr
	*/
	SBC: {
		Base:  0x0800,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    SBCI Rd, k
		Encoding  0100 KKKK dddd KKKK
	*/
	SBCI: {
		Base:  0x4000,
		Op1:   Rd_high,
		Op2:   k_8,
		Flags: 0,
	},

	/*
		Syntax    SBI A, b
		Encoding  1001 1010 AAAA Abbb
	*/
	SBI: {
		Base:  0x9A00,
		Op1:   A_5,
		Op2:   b,
		Flags: 0,
	},

	/*
		Syntax    SBIC A, b
		Encoding  1001 1001 AAAA Abbb
	*/
	SBIC: {
		Base:  0x9900,
		Op1:   A_5,
		Op2:   b,
		Flags: 0,
	},

	/*
		Syntax    SBIS A, b
		Encoding  1001 1011 AAAA Abbb
	*/
	SBIS: {
		Base:  0x9B00,
		Op1:   A_5,
		Op2:   b,
		Flags: 0,
	},

	/*
		Syntax    SBR Rd, K
		Encoding  0110 KKKK dddd KKKK
	*/
	SBR: {
		Base:  0x6000,
		Op1:   Rd,
		Op2:   k_8,
		Flags: 0,
	},

	/*
		Syntax    SBRC Rr, b
		Encoding  1111 110r rrrr 0bbb
	*/
	SBRC: {
		Base:  0xFC00,
		Op1:   Rd,
		Op2:   b,
		Flags: 0,
	},

	/*
		Syntax    SBRS Rd, b
		Encoding  1111 111r rrrr 0bbb
	*/
	SBRS: {
		Base:  0xFD00,
		Op1:   Rd,
		Op2:   b,
		Flags: 0,
	},

	/*
		Syntax    SEC
		Encoding  1001 0100 0000 1000
	*/
	SEC: {
		Base:  0x9408,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SEH
		Encoding  1001 0100 0101 1000
	*/
	SEH: {
		Base:  0x9458,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SEI
		Encoding  1001 0100 0111 1000
	*/
	SEI: {
		Base:  0x9478,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SEN
		Encoding  1001 0100 0010 1000
	*/
	SEN: {
		Base:  0x9428,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SER Rd
		Encoding  1110 1111 dddd 1111
	*/
	SER: {
		Base:  0xEF0F,
		Op1:   Rd_high,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SES
		Encoding  1001 0100 0100 1000
	*/
	SES: {
		Base:  0x9448,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SET
		Encoding  1001 0100 0110 1000
	*/
	SET: {
		Base:  0x9468,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SEV
		Encoding  1001 0100 0011 1000
	*/
	SEV: {
		Base:  0x9438,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SEZ
		Encoding  1001 0100 0001 1000
	*/
	SEZ: {
		Base:  0x9418,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    SLEEP
		Encoding  1001 0101 1000 1000
	*/
	SLEEP: {
		Base:  0x9588,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},

	/*
			Syntax    ST X+ Rr
			Encoding  1001 001d dddd 0000 kkkk kkkk kkkk kkkk

		ST: {
			Base:  0x92000000,
			Op1:   k_16,
			Op2:   R_long,
			Flags: LONG,
		},
	*/

	/*
		Syntax    STS k, Rd
		Encoding  1001 001d dddd 0000 kkkk kkkk kkkk kkkk
	*/
	STS: {
		Base:  0x92000000,
		Op1:   k_16,
		Op2:   R_long,
		Flags: LONG,
	},

	/*
		Syntax    SUB Rd, Rr
		Encoding  0001 10rd dddd rrrr
	*/
	SUB: {
		Base:  0x1800,
		Op1:   Rd,
		Op2:   Rr,
		Flags: 0,
	},

	/*
		Syntax    SUBI Rd, k
		Encoding  0101 KKKK dddd KKKK
	*/
	SUBI: {
		Base:  0x5000,
		Op1:   Rd,
		Op2:   k_8,
		Flags: 0,
	},

	/*
		Syntax    SWAP Rd
		Encoding  1001 010d dddd 0010
	*/
	SWAP: {
		Base:  0x9400,
		Op1:   Rd,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    TST Rd
		Encoding  0010 00dd dddd dddd
	*/
	TST: {
		Base:  0x2000,
		Op1:   R,
		Op2:   nil,
		Flags: 0,
	},

	/*
		Syntax    WDR
		Encoding  1001 0101 1010 1000
	*/
	WDR: {
		Base:  0x95A8,
		Op1:   nil,
		Op2:   nil,
		Flags: 0,
	},
}

var AVRe = map[Mnemonic]Instruction{
	/*
		Syntax    JMP k
		Encoding  1001 010k kkkk 110k kkkk kkkk kkkk kkkk
	*/
	JMP: {
		Base:  0x940C0000,
		Op1:   k_22,
		Op2:   nil,
		Flags: LONG,
	},
}

func (instr *Instruction) Apply1(op uint64) error {
	newBase, err := instr.Op1(instr.Base, op)

	if err != nil {
		return err
	}

	instr.Base = newBase
	return nil
}

func (instr *Instruction) Apply2(op uint64) error {
	newBase, err := instr.Op2(instr.Base, op)

	if err != nil {
		return err
	}

	instr.Base = newBase
	return nil
}

func (instr *Instruction) Encode() []byte {
	var bs []byte

	if (instr.Flags & LONG) == LONG {
		bs = make([]byte, 4)
		binary.LittleEndian.PutUint16(bs, uint16(instr.Base>>16))
		binary.LittleEndian.PutUint16(bs[2:], uint16(instr.Base))
	} else {
		bs = make([]byte, 2)
		binary.LittleEndian.PutUint16(bs, uint16(instr.Base))
	}

	return bs
}

func (instr *Instruction) IsLong() bool {
	return (instr.Flags & LONG) == LONG
}

func (instr *Instruction) IsRelative() bool {
	return (instr.Flags & RELATIVE) == RELATIVE
}

func (instr *Instruction) Print() {
	fmt.Printf("INSTR: %X\n", instr.Base)
}
