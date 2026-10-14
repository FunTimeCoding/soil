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
	s *Summaries,
) bool {
	selector, okay := c.Call.Fun.(*ast.SelectorExpr)

	if !okay {
		return false
	}

	if f := calleeOf(p, c.Call); f != nil && f.Signature().Recv() != nil &&
		!isInterfaceMethod(f) {
		if s.selfFor(f) {
			return s.chainBorrowed(p, body, selector.X)
		}

		return s.lendsFor(f)
	}

	receiver := rootIdentifier(p, selector.X)

	if receiver == nil {
		return false
	}

	return !definedIn(p, body, receiver) ||
		escapes(p, body, receiver, rootsAt) ||
		passed(p, body, receiver, rootsAt, s)
}
