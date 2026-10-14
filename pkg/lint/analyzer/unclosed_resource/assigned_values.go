package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func assignedValues(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
) []ast.Expr {
	var result []ast.Expr
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			a, okay := n.(*ast.AssignStmt)

			if !okay {
				return true
			}

			for i, l := range a.Lhs {
				i2, isIdent := l.(*ast.Ident)

				if !isIdent || (p.TypesInfo.Defs[i2] != o && p.TypesInfo.Uses[i2] != o) {
					continue
				}

				if len(a.Lhs) != len(a.Rhs) {
					result = append(result, checkedValue(a, i))

					continue
				}

				result = append(result, a.Rhs[i])
			}

			return true
		},
	)

	return result
}
