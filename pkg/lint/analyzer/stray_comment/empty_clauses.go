package stray_comment

import (
	"go/ast"
	"go/token"
)

func emptyClauses(body *ast.BlockStmt) []region {
	var result []region

	for index, clause := range body.List {
		var from token.Pos
		var size int

		switch t := clause.(type) {
		case *ast.CaseClause:
			from = t.Colon
			size = len(t.Body)
		case *ast.CommClause:
			from = t.Colon
			size = len(t.Body)
		default:
			// pass
		}

		if size != 0 {
			continue
		}

		to := body.Rbrace

		if index+1 < len(body.List) {
			to = body.List[index+1].Pos()
		}

		result = append(result, region{From: from, To: to})
	}

	return result
}
