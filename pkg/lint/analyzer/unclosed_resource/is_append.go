package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isAppend(
	p *packages.Package,
	c *ast.CallExpr,
) bool {
	i, okay := c.Fun.(*ast.Ident)

	if !okay {
		return false
	}

	_, okay = p.TypesInfo.Uses[i].(*types.Builtin)

	return okay && i.Name == "append"
}
