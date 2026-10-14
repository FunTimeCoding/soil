package unclosed_resource

import (
	"go/ast"
	"go/token"
)

func literalOf(e ast.Expr) *ast.CompositeLit {
	switch v := ast.Unparen(e).(type) {
	case *ast.UnaryExpr:
		if v.Op == token.AND {
			return literalOf(v.X)
		}
	case *ast.CompositeLit:
		return v
	}

	return nil
}
