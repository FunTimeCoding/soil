package error_answer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"slices"
)

func captures(e ast.Expr) bool {
	found := false
	ast.Inspect(
		e,
		func(n ast.Node) bool {
			call, okay := n.(*ast.CallExpr)

			if !okay {
				return !found
			}

			if x, okay := call.Fun.(*ast.SelectorExpr); okay &&
				slices.Contains(constant.CaptureCalls, x.Sel.Name) {
				found = true
			}

			return !found
		},
	)

	return found
}
