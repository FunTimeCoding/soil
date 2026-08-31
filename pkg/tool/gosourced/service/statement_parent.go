package service

import "go/ast"

func statementParent(path []ast.Node, statement ast.Node) ast.Node {
	for i, node := range path {
		if node != statement {
			continue
		}

		if i+1 >= len(path) {
			return nil
		}

		switch path[i+1].(type) {
		case *ast.BlockStmt, *ast.CaseClause, *ast.CommClause:
			return path[i+1]
		}

		return nil
	}

	return nil
}
