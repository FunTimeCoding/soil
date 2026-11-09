package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func inGuard(
	p *packages.Package,
	stack []ast.Node,
	wrappers map[types.Object]bool,
) bool {
	for i := len(stack) - 2; i >= 0; i-- {
		c, okay := stack[i].(*ast.CallExpr)

		if !okay || (!guardMethod(p, c.Fun) && !wrappers[callee(p, c.Fun)]) {
			continue
		}

		for _, a := range c.Args {
			if a == stack[i+1] {
				return true
			}
		}
	}

	return false
}
