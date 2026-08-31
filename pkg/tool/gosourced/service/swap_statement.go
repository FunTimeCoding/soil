package service

import (
	"fmt"
	"github.com/dave/dst"
)

func swapStatement(
	parent dst.Node,
	old dst.Node,
	replacement dst.Stmt,
) error {
	var list []dst.Stmt

	switch p := parent.(type) {
	case *dst.BlockStmt:
		list = p.List
	case *dst.CaseClause:
		list = p.Body
	case *dst.CommClause:
		list = p.Body
	default:
		return fmt.Errorf("unsupported statement parent")
	}

	for i, statement := range list {
		if statement == old {
			list[i] = replacement

			return nil
		}
	}

	return fmt.Errorf("statement not found in its parent")
}
