package scan

import "go/ast"

func callsPackage(
	f *ast.File,
	name string,
) bool {
	found := false
	ast.Inspect(
		f,
		func(n ast.Node) bool {
			if found {
				return false
			}

			if p, okay := calledPackage(n); okay && p == name {
				found = true
			}

			return true
		},
	)

	return found
}
