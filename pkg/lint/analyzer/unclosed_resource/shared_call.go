package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) sharedCall(
	p *packages.Package,
	c *ast.CallExpr,
) bool {
	if s == nil {
		return false
	}

	var name *ast.Ident

	switch v := ast.Unparen(c.Fun).(type) {
	case *ast.Ident:
		name = v
	case *ast.SelectorExpr:
		name = v.Sel
	default:
		return false
	}

	o, okay := p.TypesInfo.Uses[name].(*types.Var)

	if !okay || o.Pkg() == nil || o.Parent() != o.Pkg().Scope() {
		return false
	}

	return s.shared[variableKey(o)]
}
