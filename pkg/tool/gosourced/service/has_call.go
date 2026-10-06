package service

import "go/ast"

func hasCall(e ast.Expr) bool {
	found := false
	ast.Inspect(
		e,
		func(n ast.Node) bool {
			if _, okay := n.(*ast.CallExpr); okay {
				found = true
			}

			return !found
		},
	)

	return found
}
