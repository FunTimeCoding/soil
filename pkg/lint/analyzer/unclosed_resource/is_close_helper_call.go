package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isCloseHelperCall(
	p *packages.Package,
	c *ast.CallExpr,
	o *types.Var,
) bool {
	selector, okay := c.Fun.(*ast.SelectorExpr)

	if !okay || !constant.CloseHelpers[selector.Sel.Name] {
		return false
	}

	x, okay := selector.X.(*ast.Ident)

	if !okay {
		return false
	}

	name, okay := p.TypesInfo.Uses[x].(*types.PkgName)

	if !okay || name.Imported().Name() != "errors" {
		return false
	}

	return len(c.Args) > 0 && rootIdentifier(p, c.Args[0]) == o
}
