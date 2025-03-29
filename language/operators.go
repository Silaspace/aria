package language

type ArityType int

const (
	Unary  ArityType = 1
	Binary ArityType = 1 << 1
)

type Operator struct {
	BindingPower int
	Arity        ArityType
	Apply        func(uint64, uint64) uint64
}

const (
	OP_LNOT Mnemonic = "!"
	OP_BNOT Mnemonic = "~"
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
)

var Operators = map[Mnemonic]Operator{

	OP_LNOT: {
		BindingPower: 14,
		Arity:        Unary,
		Apply: func(e1 uint64, _ uint64) uint64 {
			if e1 == 0 {
				return 1
			} else {
				return 0
			}
		},
	},

	OP_BNOT: {
		BindingPower: 14,
		Arity:        Unary,
		Apply: func(e1 uint64, _ uint64) uint64 {
			return e1 ^ 0xFFFFFFFFFFFFFFFF
		},
	},

	OP_MUL: {
		BindingPower: 13,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 * e2
		},
	},

	OP_DIV: {
		BindingPower: 13,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 / e2
		},
	},

	OP_MOD: {
		BindingPower: 13,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 % e2
		},
	},

	OP_ADD: {
		BindingPower: 12,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 + e2
		},
	},

	OP_SUB: {
		BindingPower: 12,
		Arity:        Unary | Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 - e2
		},
	},

	OP_LSL: {
		BindingPower: 11,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 << e2
		},
	},

	OP_LSR: {
		BindingPower: 11,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 >> e2
		},
	},

	OP_LT: {
		BindingPower: 10,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			if e1 < e2 {
				return 1
			} else {
				return 0
			}
		},
	},

	OP_LTE: {
		BindingPower: 10,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			if e1 <= e2 {
				return 1
			} else {
				return 0
			}
		},
	},

	OP_GT: {
		BindingPower: 10,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			if e1 > e2 {
				return 1
			} else {
				return 0
			}
		},
	},

	OP_GTE: {
		BindingPower: 10,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			if e1 >= e2 {
				return 1
			} else {
				return 0
			}
		},
	},

	OP_EQ: {
		BindingPower: 9,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			if e1 == e2 {
				return 1
			} else {
				return 0
			}
		},
	},

	OP_NEQ: {
		BindingPower: 9,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			if e1 != e2 {
				return 1
			} else {
				return 0
			}
		},
	},

	OP_BAND: {
		BindingPower: 8,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 & e2
		},
	},

	OP_XOR: {
		BindingPower: 7,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 ^ e2
		},
	},

	OP_BOR: {
		BindingPower: 6,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			return e1 | e2
		},
	},

	OP_LAND: {
		BindingPower: 5,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			if e1 > 0 && e2 > 0 {
				return 1
			} else {
				return 0
			}
		},
	},

	OP_LOR: {
		BindingPower: 4,
		Arity:        Binary,
		Apply: func(e1 uint64, e2 uint64) uint64 {
			if e1 > 0 || e2 > 0 {
				return 1
			} else {
				return 0
			}
		},
	},
}

func (op *Operator) IsUnary() bool {
	return (op.Arity & Unary) == Unary
}

func (op *Operator) IsBinary() bool {
	return (op.Arity & Binary) == Binary
}
