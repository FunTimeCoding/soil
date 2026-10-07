package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_candidate"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func isBorrowed(
	p *packages.Package,
	body *ast.BlockStmt,
	c resource_candidate.Candidate,
) bool {
	selector, okay := c.Call.Fun.(*ast.SelectorExpr)

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
