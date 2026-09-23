package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func references(
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

			i, okay := n.(*ast.Ident)

			if okay && p.TypesInfo.Uses[i] == o {
				result = true
			}

			return !result
		},
	)

	return result
}
