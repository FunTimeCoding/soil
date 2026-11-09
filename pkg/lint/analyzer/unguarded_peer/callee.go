package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func callee(
	p *packages.Package,
	f ast.Expr,
) types.Object {
	var result types.Object

	switch v := f.(type) {
	case *ast.Ident:
		result = p.TypesInfo.Uses[v]
	case *ast.SelectorExpr:
		result = p.TypesInfo.Uses[v.Sel]
	case *ast.IndexExpr:
		return callee(p, v.X)
	case *ast.IndexListExpr:
		return callee(p, v.X)
	}

	if f, okay := result.(*types.Func); okay {
		return f.Origin()
	}

	return result
}
