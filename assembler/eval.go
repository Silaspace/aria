package assembler

import (
	"fmt"
	"strconv"

	"github.com/silaspace/aria/language"
	"github.com/silaspace/aria/parser"
)

func EvalArg(arg parser.Arg, symbolTable map[string]uint64, relativeInstr bool, pc uint64) language.Value {
	switch arg := arg.(type) {
	case *parser.Nil:
		return &language.Nil{}

	case *parser.ArgReg:
		regVal, err := strconv.ParseUint(arg.Value, 10, 32)

		if err != nil {
			return &language.Error{
				Value: err.Error(),
			}
		}

		return &language.Reg{
			Value: regVal,
		}

	case *parser.ArgExpr:
		val, err := EvalExpr(arg.Value, symbolTable, relativeInstr, pc)

		if err != nil {
			return &language.Error{
				Value: err.Error(),
			}
		}

		return &language.Int{
			Value: val,
		}

	default:
		return &language.Error{
			Value: "unkown arg type",
		}
	}
}

func EvalDirVal(dirval parser.DirVal, symbolTable map[string]uint64) language.Value {
	switch dirval := dirval.(type) {
	case *parser.NilDirVal:
		return &language.Nil{}

	case *parser.IdentDirVal:
		return &language.Ident{
			Value: dirval.Value,
		}

	case *parser.ImmDirVal:
		val, err := strconv.ParseUint(dirval.Value, 10, 64)

		if err != nil {
			return &language.Error{
				Value: err.Error(),
			}
		}

		return &language.Int{
			Value: val,
		}

	case *parser.ExprDirVal:
		val, err := EvalExpr(dirval.Value, symbolTable, false, 0)

		if err != nil {
			return &language.Error{
				Value: err.Error(),
			}
		}

		return &language.Int{
			Value: val,
		}

	case *parser.ExprListDirVal:

	case *parser.AssignDirVal:
		val, err := EvalExpr(dirval.Value, symbolTable, false, 0)

		if err != nil {
			return &language.Error{
				Value: err.Error(),
			}
		}

		return &language.Assignment{
			Symbol: dirval.Symbol,
			Value:  val,
		}

	case *parser.ErrorDirVal:
		return &language.Error{
			Value: dirval.Value,
		}

	default:
		return &language.Error{
			Value: fmt.Sprintf("Unknown dirval '%v'", dirval.Fmt()),
		}
	}

	return &language.Error{
		Value: "How did we get here?",
	}
}

func EvalExpr(expr parser.Expr, symbolTable map[string]uint64, relativeInstr bool, pc uint64) (uint64, error) {
	switch expr := expr.(type) {
	case *parser.Ident:
		// Return the value of pc if used in an expression
		if expr.Value == string(language.PC) {
			return pc, nil
		}

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

	case *parser.MonopExpr:
		e1, err := EvalExpr(expr.E1, symbolTable, relativeInstr, pc)

		if err != nil {
			return 0, err
		}

		return expr.Op.Apply(e1, 0), nil

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

	case *parser.FuncExpr:
		e1, err := EvalExpr(expr.E1, symbolTable, relativeInstr, pc)

		if err != nil {
			return 0, err
		}

		return expr.Func.Apply(e1), nil

	default:
		return 0, fmt.Errorf("unkown expr type")
	}
}
