package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) disposesTransferField(
	p *packages.Package,
	c *ast.CallExpr,
	t resource_transfer.Transfer,
) bool {
	f := calleeOf(p, c)

	if f == nil {
		return false
	}

	for i, a := range c.Args {
		if isTransferField(p, a, t) && s.takes(f, i) {
			return true
		}
	}

	return false
}
