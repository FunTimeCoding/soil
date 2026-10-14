package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) ownerRoot(
	p *packages.Package,
	e ast.Expr,
) *types.Var {
	for {
		c, okay := ast.Unparen(e).(*ast.CallExpr)

		if !okay {
			return rootIdentifier(p, e)
		}

		selector, isSelector := c.Fun.(*ast.SelectorExpr)

		if !isSelector || !s.selfFor(calleeOf(p, c)) {
			return nil
		}

		e = selector.X
	}
}
