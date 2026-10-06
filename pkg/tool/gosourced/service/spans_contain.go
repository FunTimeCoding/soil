package service

import "go/ast"

func spansContain(
	spans []ast.Node,
	n ast.Node,
) bool {
	for _, span := range spans {
		if span.Pos() <= n.Pos() && n.End() <= span.End() {
			return true
		}
	}

	return false
}
