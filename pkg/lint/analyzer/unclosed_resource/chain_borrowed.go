package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_candidate"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) chainBorrowed(
	p *packages.Package,
	body *ast.BlockStmt,
	receiver ast.Expr,
) bool {
	c, okay := ast.Unparen(receiver).(*ast.CallExpr)

	if !okay {
		return true
	}

	return s.isUnowned(p, body, resource_candidate.Candidate{Call: c})
}
