package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func carriers(
	p *packages.Package,
	body *ast.BlockStmt,
	e ast.Expr,
	indices map[*types.Var]int,
	visited map[*types.Var]bool,
) []fact.Forward {
	var result []fact.Forward

	switch v := ast.Unparen(e).(type) {
	case *ast.Ident:
		o, okay := p.TypesInfo.Uses[v].(*types.Var)

		if !okay || visited[o] {
			return nil
		}

		if index, parameter := indices[o]; parameter {
			return []fact.Forward{{Argument: index}}
		}

		visited[o] = true

		for _, a := range assignedValues(p, body, o) {
			if a != nil {
				result = append(
					result,
					carriers(p, body, a, indices, visited)...,
				)
			}
		}
	case *ast.UnaryExpr:
		if v.Op == token.AND {
			return carriers(p, body, v.X, indices, visited)
		}
	case *ast.CallExpr:
		f := calleeOf(p, v)

		if f == nil {
			return nil
		}

		for j, a := range v.Args {
			for _, c := range carriers(p, body, a, indices, visited) {
				result = append(
					result,
					fact.Forward{
						Argument: c.Argument,
						Conditions: append(
							append([]fact.Condition{}, c.Conditions...),
							*fact.NewWrapCondition(functionKey(f), j),
						),
					},
				)
			}
		}
	}

	return result
}
