package assembler

import (
	"fmt"
	"strconv"

	"github.com/silaspace/aria/parser"
)

func EvalArg(arg parser.Arg, symbolTable map[string]uint64, relativeInstr bool, pc uint64) (uint64, error) {

	switch arg := arg.(type) {
	case *parser.ArgReg:
		regVal, err := strconv.ParseUint(arg.Value, 10, 32)
		if err != nil {
			return 0, err
		}
		return regVal, nil

	case *parser.ArgExpr:
		val, err := EvalExpr(arg.Value, symbolTable, relativeInstr, pc)

		if err != nil {
			return 0, err
		}

		return val, nil

	default:
		return 0, fmt.Errorf("unkown arg type")
	}
}

func EvalExpr(expr parser.Expr, symbolTable map[string]uint64, relativeInstr bool, pc uint64) (uint64, error) {
	switch expr := expr.(type) {
	case *parser.Ident:
		val, exists := symbolTable[expr.Value]

		if exists && relativeInstr {
			return val - pc - 1, nil
		} else if exists {
			return val, nil
		} else {
			return 0, fmt.Errorf("identifier '%s' unknown", expr.Value)
		}

	case *parser.Literal:
		val, err := strconv.ParseUint(expr.Value, expr.Base, 64)

		if err != nil {
			return 0, err
		}

		return val, nil

	/*
		case *parser.MonopExpr:
			e1, err := EvalExpr(expr.E1)

			if err != nil {
				return 0, err
			}

			return expr.Op.Apply(e1), nil
	*/

	case *parser.BinopExpr:
		e1, err := EvalExpr(expr.E1, symbolTable, relativeInstr, pc)

		if err != nil {
			return 0, err
		}

		e2, err := EvalExpr(expr.E1, symbolTable, relativeInstr, pc)

		if err != nil {
			return 0, err
		}

		return expr.Op.Apply(e1, e2), nil

	default:
		return 0, fmt.Errorf("unkown expr type")
	}
}
