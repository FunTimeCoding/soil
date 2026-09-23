package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func anyMatches(
	p *packages.Package,
	expressions []ast.Expr,
	o *types.Var,
	match matcher,
) bool {
	for _, e := range expressions {
		if match(p, e, o) {
			return true
		}
	}

	return false
}
