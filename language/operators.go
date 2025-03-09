package language

type Operator struct {
	BindingPower int
	Apply        func(uint64, uint64) uint64
}

const (
	//OP_LNOT Mnemonic = "!"
	//OP_BNOT Mnemonic = "~"
	OP_MUL  Mnemonic = "*"
	OP_DIV  Mnemonic = "/"
	OP_MOD  Mnemonic = "%"
	OP_ADD  Mnemonic = "+"
	OP_SUB  Mnemonic = "-"
	OP_LSL  Mnemonic = "<<"
	OP_LSR  Mnemonic = ">>"
	OP_LT   Mnemonic = "<"
	OP_LTE  Mnemonic = "<="
	OP_GT   Mnemonic = ">"
	OP_GTE  Mnemonic = ">="
	OP_EQ   Mnemonic = "=="
	OP_NEQ  Mnemonic = "!="
	OP_BAND Mnemonic = "&"
	OP_XOR  Mnemonic = "^"
	OP_BOR  Mnemonic = "|"
	OP_LAND Mnemonic = "&&"
	OP_LOR  Mnemonic = "||"
	//OP_COND Mnemonic = "?"
)

var Operators = map[Mnemonic]Operator{
	/*
		OP_LNOT: {
			BindingPower: 14,
		},
		OP_BNOT: {
			BindingPower: 14,
		},
	*/
	OP_MUL: {
		BindingPower: 13,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 * e2
		},
	},
	OP_DIV: {
		BindingPower: 13,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 / e2
		},
	},
	OP_MOD: {
		BindingPower: 13,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 % e2
		},
	},
	OP_ADD: {
		BindingPower: 12,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 + e2
		},
	},
	OP_SUB: {
		BindingPower: 12,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 - e2
		},
	},
	OP_LSL: {
		BindingPower: 11,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 << e2
		},
	},
	OP_LSR: {
		BindingPower: 11,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 >> e2
		},
	},
	OP_LT: {
		BindingPower: 10,
		Apply: func(e1, e2 uint64) uint64 {
			if e1 < e2 {
				return 1
			} else {
				return 0
			}
		},
	},
	OP_LTE: {
		BindingPower: 10,
		Apply: func(e1, e2 uint64) uint64 {
			if e1 <= e2 {
				return 1
			} else {
				return 0
			}
		},
	},
	OP_GT: {
		BindingPower: 10,
		Apply: func(e1, e2 uint64) uint64 {
			if e1 > e2 {
				return 1
			} else {
				return 0
			}
		},
	},
	OP_GTE: {
		BindingPower: 10,
		Apply: func(e1, e2 uint64) uint64 {
			if e1 >= e2 {
				return 1
			} else {
				return 0
			}
		},
	},
	OP_EQ: {
		BindingPower: 9,
		Apply: func(e1, e2 uint64) uint64 {
			if e1 == e2 {
				return 1
			} else {
				return 0
			}
		},
	},
	OP_NEQ: {
		BindingPower: 9,
		Apply: func(e1, e2 uint64) uint64 {
			if e1 != e2 {
				return 1
			} else {
				return 0
			}
		},
	},
	OP_BAND: {
		BindingPower: 8,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 & e2
		},
	},
	OP_XOR: {
		BindingPower: 7,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 ^ e2
		},
	},
	OP_BOR: {
		BindingPower: 6,
		Apply: func(e1, e2 uint64) uint64 {
			return e1 | e2
		},
	},
	OP_LAND: {
		BindingPower: 5,
		Apply: func(e1, e2 uint64) uint64 {
			if e1 > 0 && e2 > 0 {
				return 1
			} else {
				return 0
			}
		},
	},
	OP_LOR: {
		BindingPower: 4,
		Apply: func(e1, e2 uint64) uint64 {
			if e1 > 0 || e2 > 0 {
				return 1
			} else {
				return 0
			}
		},
	},
	/*
		OP_COND: {
			BindingPower: 3,
		},
	*/
}
