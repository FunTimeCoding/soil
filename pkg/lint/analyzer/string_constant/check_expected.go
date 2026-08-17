package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkExpected(
	p *packages.Package,
	results *output.Results,
	expected ast.Expr,
) {
	ast.Inspect(
		expected,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.IndexExpr:
				checkExpected(p, results, v.X)

				return false
			case *ast.SliceExpr:
				checkExpected(p, results, v.X)

				return false
			case *ast.Ident:
				checkConstant(p, results, v)
			}

			return true
		},
	)
}
