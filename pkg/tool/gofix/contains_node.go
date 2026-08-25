package gofix

import "github.com/dave/dst"

func containsNode(
	parent dst.Node,
	target dst.Node,
) bool {
	if parent == target {
		return true
	}

	switch n := parent.(type) {
	case *dst.UnaryExpr:
		return containsNode(n.X, target)
	case *dst.SelectorExpr:
		return containsNode(n.X, target)
	case *dst.CallExpr:
		return containsNode(n.Fun, target)
	case *dst.StarExpr:
		return containsNode(n.X, target)
	case *dst.ParenExpr:
		return containsNode(n.X, target)
	}

	return false
}
