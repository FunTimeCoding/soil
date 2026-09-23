package unclosed_resource

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func isBorrowed(
	p *packages.Package,
	body *ast.BlockStmt,
	c candidate,
) bool {
	selector, okay := c.call.Fun.(*ast.SelectorExpr)

	if !okay {
		return false
	}

	receiver := rootIdentifier(p, selector.X)

	if receiver == nil {
		return false
	}

	return !definedIn(p, body, receiver) ||
		escapes(p, body, receiver, rootsAt)
}
