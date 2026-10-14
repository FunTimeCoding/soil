package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func releasesIn(
	p *packages.Package,
	body *ast.BlockStmt,
	r *types.Var,
) []fact.Release {
	var result []fact.Release
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.CallExpr:
				result = append(result, callRelease(p, v, r)...)
			case *ast.RangeStmt:
				if field := fieldOfReceiver(p, v.X, r); field != "" {
					result = append(result, fact.Release{Field: field})
				}
			}

			return true
		},
	)

	return result
}
