package language

type Function struct {
	Apply func(uint64) uint64
}

const (
	FUNC_LOW   Mnemonic = "low"
	FUNC_HIGH  Mnemonic = "high"
	FUNC_BYTE2 Mnemonic = "byte2"
	FUNC_BYTE3 Mnemonic = "byte3"
	FUNC_BYTE4 Mnemonic = "byte4"
	FUNC_LWRD  Mnemonic = "lwrd"
	FUNC_HWRD  Mnemonic = "hwrd"
	FUNC_PAGE  Mnemonic = "page"
	FUNC_EXP2  Mnemonic = "exp2"
	FUNC_LOG2  Mnemonic = "log2"
)

var Functions = map[Mnemonic]Function{
	FUNC_LOW: {
		Apply: func(e uint64) uint64 {
			return e & 0xFF
		},
	},

	FUNC_HIGH: {
		Apply: func(e uint64) uint64 {
			return (e & 0xFF00) >> 8
		},
	},

	FUNC_BYTE2: {
		Apply: func(e uint64) uint64 {
			return (e & 0xFF00) >> 8
		},
	},

	FUNC_BYTE3: {
		Apply: func(e uint64) uint64 {
			return (e & 0xFF0000) >> 16
		},
	},

	FUNC_BYTE4: {
		Apply: func(e uint64) uint64 {
			return (e & 0xFF000000) >> 24
		},
	},

	FUNC_LWRD: {
		Apply: func(e uint64) uint64 {
			return e & 0xFFFF
		},
	},

	FUNC_HWRD: {
		Apply: func(e uint64) uint64 {
			return (e & 0xFFFF0000) >> 16
		},
	},

	FUNC_PAGE: {
		Apply: func(e uint64) uint64 {
			return (e & 0xF0000) >> 16
		},
	},

	FUNC_EXP2: {
		Apply: func(e uint64) uint64 {
			return exp(2, e)
		},
	},

	FUNC_LOG2: {
		Apply: func(e uint64) uint64 {
			return intlog(2, e)
		},
	},
}

func exp(x uint64, y uint64) uint64 {
	// x ^ 2y = (x*x) ^ y
	// x ^ (2y+1) = (x ^ 2y) * x

	var z uint64 = 1

	for {
		if y == 0 {
			return z
		} else if (y % 2) == 0 {
			x = x * x
			y = y / 2
		} else {
			z = z * x
			y = y - 1
		}
	}
}

func intlog(a uint64, b uint64) uint64 {
	var c uint64 = 1
	var d uint64 = 0

	for {
		c2 := c * a
		if c2 > b {
			return c
		} else {
			d = d + 1
			c = c2
		}
	}
}
