package language

import (
	"fmt"

	"github.com/silaspace/aria/device"
)

type Mnemonic string

type KeywordType int

const (
	IDENT KeywordType = 0
	INSTR KeywordType = 1
	DIR   KeywordType = 2
	FUNC  KeywordType = 3
	REG   KeywordType = 4
)

func Exists(key string) KeywordType {

	mn := Mnemonic(key)

	if _, exists := AVR[mn]; exists {
		return INSTR
	}

	if _, exists := AVRe[mn]; exists {
		return INSTR
	}

	if _, exists := Directives[mn]; exists {
		return DIR
	}

	if _, exists := Functions[mn]; exists {
		return FUNC
	}

	if _, exists := Registers[mn]; exists {
		return REG
	}

	return IDENT
}

func GetInstr(key string, dev *device.Device) (Instruction, error) {
	mn := Mnemonic(key)

	var instr Instruction
	var exist bool

	/*
		Core specific instructions have priority over core
		instructions. This allows cores such as AVRrc to have
		the same mnemonic but encode differently to save space.
	*/

	switch dev.DeviceCore {
	case device.Nil:
		if instr, exist = AVR_core[mn]; exist {
			return instr, nil
		}

		return instr, fmt.Errorf("device core not specified; '%v' does not exist", key)

	case device.AVR:
		if instr, exist = AVR[mn]; exist {
			return instr, nil
		}

		if instr, exist = AVR_core[mn]; exist {
			return instr, nil
		}

		return instr, fmt.Errorf("'%v' does not exist in the AVR instruction set", key)

	case device.AVRe:
		if instr, exist = AVRe[mn]; exist {
			return instr, nil
		}

		if instr, exist = AVR[mn]; exist {
			return instr, nil
		}

		if instr, exist = AVR_core[mn]; exist {
			return instr, nil
		}

		return instr, fmt.Errorf("'%v' does not exist in the AVRe instruction set", key)

	default:
		return instr, fmt.Errorf("device core '%v' not implemented", dev.DeviceCore)
	}
}

func GetOp(key string) (Operator, error) {
	mn := Mnemonic(key)
	op, exists := Operators[mn]

	if !exists {
		return op, fmt.Errorf("operator '%v' does not exist", key)
	}

	return op, nil
}

func GetDir(key string) (Directive, error) {
	mn := Mnemonic(key)
	dir, exists := Directives[mn]

	if !exists {
		return dir, fmt.Errorf("directive '%v' does not exist", key)
	}

	return dir, nil
}

func GetFunc(key string) (Function, error) {
	mn := Mnemonic(key)
	function, exists := Functions[mn]

	if !exists {
		return function, fmt.Errorf("directive '%v' does not exist", key)
	}

	return function, nil
}

func IsPC(key string) bool {
	mn := Mnemonic(key)
	return mn == PC
}

func IsAdd(key string) bool {
	mn := Mnemonic(key)
	return mn == OP_ADD
}

func IsSub(key string) bool {
	mn := Mnemonic(key)
	return mn == OP_SUB
}
