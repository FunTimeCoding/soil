package parse

import "go/ast"

func FindMethods(
	f *ast.File,
	name string,
) []*ast.CallExpr {
	var result []*ast.CallExpr
	ast.Inspect(
		f,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			s, okay := c.Fun.(*ast.SelectorExpr)

			if okay && s.Sel.Name == name {
				result = append(result, c)
			}

			return true
		},
	)

	return result
}
