package forwarding_function

import "go/ast"

func plainCallee(e ast.Expr) bool {
	result := true
	ast.Inspect(
		e,
		func(n ast.Node) bool {
			if _, okay := n.(*ast.CallExpr); okay {
				result = false
			}

			return result
		},
	)

	return result
}
