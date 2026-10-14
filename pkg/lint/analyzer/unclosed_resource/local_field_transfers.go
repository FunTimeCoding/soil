package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func localFieldTransfers(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
) []resource_transfer.Transfer {
	var result []resource_transfer.Transfer
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			a, okay := n.(*ast.AssignStmt)

			if !okay || len(a.Lhs) != len(a.Rhs) {
				return true
			}

			for i, l := range a.Lhs {
				s, isSelector := l.(*ast.SelectorExpr)

				if !isSelector || rootIdentifier(p, s.X) != o {
					continue
				}

				holder := selectedHolder(p, s)
				field, isField := p.TypesInfo.Selections[s].Obj().(*types.Var)

				if holder == nil || !isField {
					continue
				}

				result = append(
					result,
					resource_transfer.Transfer{
						Holder: holder,
						Field:  field,
						Value:  a.Rhs[i],
					},
				)
			}

			return true
		},
	)

	return result
}
