package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isNil(
	p *packages.Package,
	e ast.Expr,
) bool {
	i, okay := ast.Unparen(e).(*ast.Ident)

	if !okay {
		return false
	}

	_, result := p.TypesInfo.Uses[i].(*types.Nil)

	return result
}
