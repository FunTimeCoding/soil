package service

import "github.com/dave/dst"

func removeBlockStatement(
	block *dst.BlockStmt,
	statement dst.Stmt,
) bool {
	for i, s := range block.List {
		if s == statement {
			block.List = append(block.List[:i], block.List[i+1:]...)

			return true
		}
	}

	return false
}
