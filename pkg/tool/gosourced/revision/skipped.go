package revision

import "go/ast"

func skipped(n ast.Node) map[*ast.Ident]bool {
	result := map[*ast.Ident]bool{}
	mark := func(identifiers ...*ast.Ident) {
		for _, i := range identifiers {
			if i != nil {
				result[i] = true
			}
		}
	}
	ast.Inspect(
		n,
		func(c ast.Node) bool {
			switch c := c.(type) {
			case *ast.SelectorExpr:
				mark(c.Sel)
			case *ast.Field:
				mark(c.Names...)
			case *ast.CompositeLit:
				mark(compositeKeys(c)...)
			case *ast.LabeledStmt:
				mark(c.Label)
			case *ast.BranchStmt:
				mark(c.Label)
			case *ast.FuncDecl:
				mark(c.Name)
			case *ast.TypeSpec:
				mark(c.Name)
			case *ast.ValueSpec:
				mark(c.Names...)
			}

			return true
		},
	)

	return result
}
