package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type matcher func(
	p *packages.Package,
	e ast.Expr,
	o *types.Var,
) bool
