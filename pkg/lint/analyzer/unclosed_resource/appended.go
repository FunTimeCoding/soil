package unclosed_resource

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func appended(
	p *packages.Package,
	values []ast.Expr,
) []ast.Expr {
	result := make([]ast.Expr, 0, len(values))

	for _, v := range values {
		if c, okay := ast.Unparen(v).(*ast.CallExpr); okay && isAppend(p, c) {
			result = append(result, appended(p, c.Args[1:])...)

			continue
		}

		result = append(result, v)
	}

	return result
}
