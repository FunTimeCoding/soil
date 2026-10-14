package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) disposesArgument(
	p *packages.Package,
	c *ast.CallExpr,
	o *types.Var,
) bool {
	if l, invoked := ast.Unparen(c.Fun).(*ast.FuncLit); invoked {
		return s.invokedDisposes(p, l, c, o)
	}

	f := calleeOf(p, c)

	if f == nil {
		return false
	}

	for i, a := range c.Args {
		if rootIdentifier(p, a) == o && s.takes(f, i) {
			return true
		}
	}

	return false
}
