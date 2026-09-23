package unclosed_resource

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func directArranges(
	p *packages.Package,
	body *ast.BlockStmt,
) bool {
	var result bool
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			if result {
				return false
			}

			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			subject := isRegistrarCall(c)

			if subject == nil {
				return true
			}

			o := rootIdentifier(p, subject)

			if o != nil && flowsToResult(p, body, o) {
				result = true

				return false
			}

			return true
		},
	)

	return result
}
