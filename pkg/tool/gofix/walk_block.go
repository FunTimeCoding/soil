package gofix

import "github.com/dave/dst"

func walkBlock(
	block *dst.BlockStmt,
	parentLit *dst.CompositeLit,
	walk func(
		dst.Node,
		*dst.CompositeLit,
		int,
	),
) {
	for _, s := range block.List {
		walkStatement(s, parentLit, walk)
	}
}
