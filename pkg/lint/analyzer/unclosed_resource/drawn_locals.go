package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func drawnLocals(
	p *packages.Package,
	body *ast.BlockStmt,
	parameter *types.Var,
) map[*types.Var]string {
	result := map[*types.Var]string{}
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			a, okay := n.(*ast.AssignStmt)

			if !okay || len(a.Rhs) != 1 || len(a.Lhs) == 0 {
				return true
			}

			c, okay := ast.Unparen(a.Rhs[0]).(*ast.CallExpr)

			if !okay {
				return true
			}

			field, okay := factoryField(p, c, parameter)

			if !okay {
				return true
			}

			if i, okay := a.Lhs[0].(*ast.Ident); okay {
				if v, okay := p.TypesInfo.ObjectOf(i).(*types.Var); okay {
					result[v] = field
				}
			}

			return true
		},
	)

	return result
}
