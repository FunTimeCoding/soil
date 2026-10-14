package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func closes(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
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
				result = s.isReleasing(p, v) && s.ownerRoot(p, v.X) == o
			case *ast.CallExpr:
				result = s.disposesArgument(p, v, o)
			}

			return !result
		},
	)

	return result
}
