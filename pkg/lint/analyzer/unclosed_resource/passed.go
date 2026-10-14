package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func passed(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
	match matcher,
	s *Summaries,
) bool {
	var result bool
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			if result {
				return false
			}

			if c, okay := n.(*ast.CallExpr); okay {
				result = !s.disposesArgument(p, c, o) &&
					anyMatches(p, c.Args, o, match)
			}

			return !result
		},
	)

	return result
}
