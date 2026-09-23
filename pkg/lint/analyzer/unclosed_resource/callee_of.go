package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func calleeOf(
	p *packages.Package,
	c *ast.CallExpr,
) *types.Func {
	var name *ast.Ident

	switch v := c.Fun.(type) {
	case *ast.Ident:
		name = v
	case *ast.SelectorExpr:
		name = v.Sel
	default:
		return nil
	}

	result, _ := p.TypesInfo.Uses[name].(*types.Func)

	return result
}
