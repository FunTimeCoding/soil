package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func rootsAt(
	p *packages.Package,
	e ast.Expr,
	o *types.Var,
) bool {
	return rootIdentifier(p, e) == o
}
