package language

type Directive struct{}

const (
	DIR_DEVICE Mnemonic = "device"
	DIR_EQU    Mnemonic = "equ"
)

var Directives = map[Mnemonic]Directive{
	DIR_DEVICE: {},
	DIR_EQU:    {},
}
