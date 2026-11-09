package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func enclosing(
	p *packages.Package,
	stack []ast.Node,
) types.Object {
	for i := len(stack) - 1; i >= 0; i-- {
		if f, okay := stack[i].(*ast.FuncDecl); okay {
			return p.TypesInfo.Defs[f.Name]
		}
	}

	return nil
}
