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

	return IDENT
}

func GetInstr(key string, dev *device.Device) (Instruction, error) {
	mn := Mnemonic(key)

	var instr Instruction
	var exist bool

	switch dev.DeviceCore {
	case device.AVR:
		break
	case device.AVRe:
		instr, exist = AVRe[mn]
	default:
		return instr, fmt.Errorf("device core '%v' not implemented", dev.DeviceCore)
	}

	if exist {
		return instr, nil
	}

	instr, exist = AVR[mn]

	if exist {
		return instr, nil
	}

	return instr, fmt.Errorf("instruction '%v' does not exist", key)
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
