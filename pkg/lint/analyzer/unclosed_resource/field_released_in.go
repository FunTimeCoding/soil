package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func fieldReleasedIn(
	p *packages.Package,
	body *ast.BlockStmt,
	t resource_transfer.Transfer,
	s *Summaries,
) bool {
	var result bool
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			if result {
				return false
			}

			switch v := n.(type) {
			case *ast.SelectorExpr:
				result = s.isReleasing(p, v) && isTransferField(p, v.X, t)
			case *ast.CallExpr:
				result = s.disposesTransferField(p, v, t)
			}

			return !result
		},
	)

	return result
}
