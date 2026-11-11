package pattern

import "go/ast"

func siteSelection(
	path []ast.Node,
	expression bool,
) (ast.Node, ast.Node) {
	if !expression {
		return siteNode(path), anchorNode(path)
	}

	anchor := anchorNode(path)

	if call, okay := anchor.(*ast.CallExpr); okay {
		return call, call.Fun
	}

	return anchor, anchor
}
