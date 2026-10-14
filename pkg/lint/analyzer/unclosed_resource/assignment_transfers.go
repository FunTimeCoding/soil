package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func assignmentTransfers(
	p *packages.Package,
	a *ast.AssignStmt,
) []resource_transfer.Transfer {
	if len(a.Lhs) != len(a.Rhs) {
		return nil
	}

	var result []resource_transfer.Transfer

	for i, l := range a.Lhs {
		target := l
		container := false

		if index, isIndex := l.(*ast.IndexExpr); isIndex {
			target = index.X
			container = true
		}

		s, okay := target.(*ast.SelectorExpr)

		if !okay {
			continue
		}

		holder := selectedHolder(p, s)

		if holder == nil {
			continue
		}

		field, _ := p.TypesInfo.Selections[s].Obj().(*types.Var)
		values := []ast.Expr{a.Rhs[i]}

		if c, isCall := a.Rhs[i].(*ast.CallExpr); isCall && !container &&
			isAppend(p, c) && len(c.Args) > 1 {
			values = c.Args[1:]
			container = true
		}

		for _, v := range values {
			result = append(
				result,
				resource_transfer.Transfer{
					Holder:    holder,
					Field:     field,
					Container: container,
					Value:     v,
				},
			)
		}
	}

	return result
}
