package language

type Function struct{}

const (
	FUNC_LOW     Mnemonic = "low"
	FUNC_HIGH    Mnemonic = "high"
	FUNC_BYTE2   Mnemonic = "byte2"
	FUNC_BYTE3   Mnemonic = "byte3"
	FUNC_BYTE4   Mnemonic = "byte4"
	FUNC_LWRD    Mnemonic = "lwrd"
	FUNC_HWRD    Mnemonic = "hwrd"
	FUNC_PAGE    Mnemonic = "page"
	FUNC_EXP2    Mnemonic = "exp2"
	FUNC_LOG2    Mnemonic = "log2"
	FUNC_INT     Mnemonic = "int"
	FUNC_FRAC    Mnemonic = "frac"
	FUNC_Q7      Mnemonic = "q7"
	FUNC_Q15     Mnemonic = "q15"
	FUNC_ABS     Mnemonic = "abs"
	FUNC_DEFINED Mnemonic = "defined"
	FUNC_STRLEN  Mnemonic = "strlen"
)

var Functions = map[Mnemonic]Function{
	FUNC_LOW:     {},
	FUNC_HIGH:    {},
	FUNC_BYTE2:   {},
	FUNC_BYTE3:   {},
	FUNC_BYTE4:   {},
	FUNC_LWRD:    {},
	FUNC_HWRD:    {},
	FUNC_PAGE:    {},
	FUNC_EXP2:    {},
	FUNC_LOG2:    {},
	FUNC_INT:     {},
	FUNC_FRAC:    {},
	FUNC_Q7:      {},
	FUNC_Q15:     {},
	FUNC_ABS:     {},
	FUNC_DEFINED: {},
	FUNC_STRLEN:  {},
}
