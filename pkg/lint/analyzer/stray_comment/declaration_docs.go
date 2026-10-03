package stray_comment

import "go/ast"

func declarationDocs(file *ast.File) map[*ast.CommentGroup]bool {
	result := map[*ast.CommentGroup]bool{}
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			switch t := n.(type) {
			case *ast.FuncDecl:
				result[t.Doc] = true
			case *ast.GenDecl:
				result[t.Doc] = true
			case *ast.TypeSpec:
				result[t.Doc] = true
			case *ast.ValueSpec:
				result[t.Doc] = true
			}

			return true
		},
	)
	delete(result, nil)

	return result
}
