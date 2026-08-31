package service

import "go/ast"

func anchorNode(path []ast.Node) ast.Node {
	result := path[0]

	for i, node := range path {
		switch n := node.(type) {
		case *ast.Ident, *ast.SelectorExpr:
			result = node

			continue
		case *ast.CallExpr:
			if i > 0 && n.Fun == path[i-1] {
				result = node
			}
		}

		break
	}

	return result
}
