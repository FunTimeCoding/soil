package census

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func called(
	p *packages.Package,
	c *ast.CallExpr,
) *types.Func {
	var i *ast.Ident

	switch x := c.Fun.(type) {
	case *ast.SelectorExpr:
		i = x.Sel
	case *ast.Ident:
		i = x
	default:
		return nil
	}

	f, _ := p.TypesInfo.Uses[i].(*types.Func)

	return f
}
