package unclosed_resource

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func storedValues(
	p *packages.Package,
	target ast.Expr,
	value ast.Expr,
) []ast.Expr {
	c, okay := ast.Unparen(value).(*ast.CallExpr)

	if !okay || !isAppend(p, c) || len(c.Args) < 2 {
		return []ast.Expr{value}
	}

	if s, isSelector := c.Args[0].(*ast.SelectorExpr); isSelector {
		if t, sameSelector := target.(*ast.SelectorExpr); sameSelector &&
			t.Sel.Name == s.Sel.Name {
			return c.Args[1:]
		}
	}

	return []ast.Expr{value}
}
