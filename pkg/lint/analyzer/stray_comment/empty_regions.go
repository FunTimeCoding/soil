package stray_comment

import "go/ast"

func emptyRegions(file *ast.File) []Region {
	var result []Region
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			switch t := n.(type) {
			case *ast.BlockStmt:
				if len(t.List) == 0 {
					result = append(
						result,
						Region{Start: t.Lbrace, End: t.Rbrace},
					)
				}
			case *ast.InterfaceType:
				if len(t.Methods.List) == 0 {
					result = append(
						result,
						Region{Start: t.Methods.Opening, End: t.Methods.Closing},
					)
				}
			case *ast.SwitchStmt:
				result = append(result, emptyClauses(t.Body)...)
			case *ast.TypeSwitchStmt:
				result = append(result, emptyClauses(t.Body)...)
			case *ast.SelectStmt:
				result = append(result, emptyClauses(t.Body)...)
			}

			return true
		},
	)

	return result
}
