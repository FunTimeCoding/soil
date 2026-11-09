package unguarded_peer

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func callsGuard(
	p *packages.Package,
	body *ast.BlockStmt,
) bool {
	result := false
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			if c, okay := n.(*ast.CallExpr); okay && guardMethod(p, c.Fun) {
				result = true
			}

			return !result
		},
	)

	return result
}
