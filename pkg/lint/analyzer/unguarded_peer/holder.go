package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func holder(
	p *packages.Package,
	x ast.Expr,
) types.Object {
	switch v := x.(type) {
	case *ast.Ident:
		return p.TypesInfo.Uses[v]
	case *ast.SelectorExpr:
		s := p.TypesInfo.Selections[v]

		if s != nil && s.Kind() == types.FieldVal {
			return s.Obj()
		}
	}

	return nil
}
