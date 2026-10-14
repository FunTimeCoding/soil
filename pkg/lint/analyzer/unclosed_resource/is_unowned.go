package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_candidate"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) isUnowned(
	p *packages.Package,
	body *ast.BlockStmt,
	c resource_candidate.Candidate,
) bool {
	f := calleeOf(p, c.Call)

	return isBorrowed(p, body, c, s) ||
		s.arrangesFor(f) ||
		s.sharedCall(p, c.Call) ||
		(s.lendsFor(f) && !s.selfFor(f)) ||
		s.spentFor(f)
}
