package gofix

import "github.com/dave/dst"

func structFieldPadding(
	node *dst.CompositeLit,
	parentLit *dst.CompositeLit,
) int {
	if parentLit == nil {
		return 0
	}

	return fieldPadding(node, parentLit)
}
