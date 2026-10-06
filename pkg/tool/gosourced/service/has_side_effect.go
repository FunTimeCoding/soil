package service

import (
	"go/ast"
	"go/token"
	"go/types"
)

func hasSideEffect(
	information *types.Info,
	e ast.Expr,
) bool {
	result := false
	ast.Inspect(
		e,
		func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.UnaryExpr:
				if x.Op == token.ARROW {
					result = true
				}
			case *ast.CallExpr:
				if !information.Types[x.Fun].IsType() {
					result = true
				}
			}

			return !result
		},
	)

	return result
}
