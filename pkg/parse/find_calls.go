package parse

import "go/ast"

func FindCalls(
	f *ast.File,
	p string,
	function string,
) []*ast.CallExpr {
	var result []*ast.CallExpr
	ast.Inspect(
		f,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if okay && matchesCall(c, p, function) {
				result = append(result, c)
			}

			return true
		},
	)

	return result
}
