package language

import "fmt"

type Directive struct {
	Execute func(Assembler, Value) error
}

type Assembler interface {
	AddSymbol(string, uint64) error
	SetDevice(string) error
}

const (
	DIR_DEVICE Mnemonic = "device"
	DIR_EQU    Mnemonic = "equ"
)

var Directives = map[Mnemonic]Directive{
	DIR_DEVICE: {
		Execute: func(a Assembler, v Value) error {
			switch v := v.(type) {
			case *Ident:
				err := a.SetDevice(v.Value)

				if err != nil {
					return err
				}

				return nil

			default:
				return fmt.Errorf("expected device, got '%v'", v)
			}
		},
	},
	DIR_EQU: {
		Execute: func(a Assembler, v Value) error {
			switch v := v.(type) {
			case *Assignment:
				err := a.AddSymbol(v.Symbol, v.Value)

				if err != nil {
					return err
				}

				return nil

			default:
				return fmt.Errorf("expected assignment, got '%v'", v)
			}
		},
	},
}
