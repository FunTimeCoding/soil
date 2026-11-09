package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func parameter(
	p *packages.Package,
	e ast.Expr,
	indices map[*types.Var]int,
) *types.Var {
	switch v := ast.Unparen(e).(type) {
	case *ast.Ident:
		o, _ := p.TypesInfo.Uses[v].(*types.Var)

		if _, okay := indices[o]; okay {
			return o
		}
	case *ast.SelectorExpr:
		o := parameter(p, v.X, indices)

		if o != nil && isResponse(o.Type()) && v.Sel.Name == "Body" {
			return o
		}
	}

	return nil
}
