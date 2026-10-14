package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func fieldOfReceiver(
	p *packages.Package,
	e ast.Expr,
	r *types.Var,
) string {
	for {
		switch v := e.(type) {
		case *ast.SelectorExpr:
			if isReceiver(p, v.X, r) {
				s := p.TypesInfo.Selections[v]

				if s == nil || s.Kind() != types.FieldVal {
					return ""
				}

				return fieldName(r.Type(), s.Index()[0])
			}

			e = v.X
		case *ast.ParenExpr:
			e = v.X
		case *ast.StarExpr:
			e = v.X
		case *ast.IndexExpr:
			e = v.X
		default:
			return ""
		}
	}
}
