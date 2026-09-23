package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func carriesFrom(
	p *packages.Package,
	e ast.Expr,
	o *types.Var,
) bool {
	if rootIdentifier(p, e) != o {
		return false
	}

	return carriesObligation(p.TypesInfo.TypeOf(e))
}
