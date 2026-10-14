package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isReceiver(
	p *packages.Package,
	e ast.Expr,
	r *types.Var,
) bool {
	for {
		switch v := e.(type) {
		case *ast.ParenExpr:
			e = v.X
		case *ast.Ident:
			return r != nil && p.TypesInfo.Uses[v] == r
		default:
			return false
		}
	}
}
