package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func factoryCarriers(
	p *packages.Package,
	body *ast.BlockStmt,
	e ast.Expr,
	indices map[*types.Var]int,
) []fact.Forward {
	l, okay := ast.Unparen(e).(*ast.FuncLit)

	if !okay {
		return nil
	}

	var result []fact.Forward
	ast.Inspect(
		l.Body,
		func(n ast.Node) bool {
			if r, okay := n.(*ast.ReturnStmt); okay {
				for _, v := range r.Results {
					result = append(
						result,
						carriers(p, body, v, indices, map[*types.Var]bool{})...,
					)
				}
			}

			return true
		},
	)

	return result
}
