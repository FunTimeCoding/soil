package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func flowsToResult(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
) bool {
	var result bool
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			if result {
				return false
			}

			r, okay := n.(*ast.ReturnStmt)

			if !okay {
				return true
			}

			for _, e := range r.Results {
				if expressionHolds(p, e, o) {
					result = true

					return false
				}
			}

			return true
		},
	)

	return result
}
