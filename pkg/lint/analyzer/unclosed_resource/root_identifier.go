package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func rootIdentifier(
	p *packages.Package,
	e ast.Expr,
) *types.Var {
	for {
		switch v := e.(type) {
		case *ast.Ident:
			result, okay := p.TypesInfo.Uses[v].(*types.Var)

			if !okay {
				return nil
			}

			return result
		case *ast.SelectorExpr:
			e = v.X
		case *ast.ParenExpr:
			e = v.X
		case *ast.StarExpr:
			e = v.X
		case *ast.IndexExpr:
			e = v.X
		default:
			return nil
		}
	}
}
