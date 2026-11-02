package unclosed_resource

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) freshRoot(
	p *packages.Package,
	c *ast.CallExpr,
) bool {
	for {
		f := calleeOf(p, c)

		if f == nil {
			return false
		}

		if f.Signature().Recv() == nil {
			return true
		}

		if !s.selfFor(f) {
			return false
		}

		selector, okay := ast.Unparen(c.Fun).(*ast.SelectorExpr)

		if !okay {
			return false
		}

		inner, okay := ast.Unparen(selector.X).(*ast.CallExpr)

		if !okay {
			return false
		}

		c = inner
	}
}
