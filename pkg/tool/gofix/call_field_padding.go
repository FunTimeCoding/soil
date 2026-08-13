package gofix

import "github.com/dave/dst"

func callFieldPadding(
	node *dst.CallExpr,
	parentLit *dst.CompositeLit,
) int {
	if parentLit == nil {
		return 0
	}

	return fieldPadding(node, parentLit)
}
