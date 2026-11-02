package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func factoryField(
	p *packages.Package,
	c *ast.CallExpr,
	parameter *types.Var,
) (string, bool) {
	switch v := ast.Unparen(c.Fun).(type) {
	case *ast.Ident:
		return "", p.TypesInfo.Uses[v] == parameter
	case *ast.SelectorExpr:
		s, okay := p.TypesInfo.Selections[v]

		if !okay || s.Kind() != types.FieldVal {
			return "", false
		}

		i, okay := ast.Unparen(v.X).(*ast.Ident)

		return v.Sel.Name, okay && p.TypesInfo.Uses[i] == parameter
	}

	return "", false
}
