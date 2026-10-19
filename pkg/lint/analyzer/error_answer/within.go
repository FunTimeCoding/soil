package error_answer

import "go/ast"

func within(
	nodes []ast.Node,
	e ast.Node,
) bool {
	for _, n := range nodes {
		if n.Pos() <= e.Pos() && e.End() <= n.End() {
			return true
		}
	}

	return false
}
