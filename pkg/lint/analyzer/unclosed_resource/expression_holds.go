package unclosed_resource

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func expressionHolds(
	p *packages.Package,
	e ast.Expr,
	o *types.Var,
) bool {
	switch v := e.(type) {
	case *ast.UnaryExpr:
		return v.Op == token.AND && expressionHolds(p, v.X, o)
	case *ast.CompositeLit:
		for _, element := range elementValues(v.Elts) {
			if expressionHolds(p, element, o) {
				return true
			}
		}

		return false
	}

	return rootIdentifier(p, e) == o
}
