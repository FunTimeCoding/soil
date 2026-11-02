package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func returnedRoot(
	p *packages.Package,
	e ast.Expr,
	returned map[types.Object]bool,
) bool {
	r := rootIdentifier(p, e)

	return r != nil && returned[r]
}
