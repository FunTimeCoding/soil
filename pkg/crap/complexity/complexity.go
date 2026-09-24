package complexity

import (
	"go/ast"
	"go/token"
)

func Complexity(declaration *ast.FuncDecl) int {
	result := 1
	ast.Inspect(
		declaration,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt:
				result++
			case *ast.CaseClause:
				if v.List != nil {
					result++
				}
			case *ast.CommClause:
				if v.Comm != nil {
					result++
				}
			case *ast.BinaryExpr:
				if v.Op == token.LAND || v.Op == token.LOR {
					result++
				}
			}

			return true
		},
	)

	return result
}
