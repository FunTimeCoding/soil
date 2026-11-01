package revision

import (
	"go/ast"
	"go/token"
)

func localNames(n ast.Node) map[string]bool {
	result := map[string]bool{}
	add := func(expressions ...ast.Expr) {
		for _, x := range expressions {
			if i, isIdent := x.(*ast.Ident); isIdent {
				result[i.Name] = true
			}
		}
	}
	ast.Inspect(
		n,
		func(c ast.Node) bool {
			switch c := c.(type) {
			case *ast.Field:
				for _, name := range c.Names {
					result[name.Name] = true
				}
			case *ast.AssignStmt:
				if c.Tok == token.DEFINE {
					add(c.Lhs...)
				}
			case *ast.RangeStmt:
				if c.Tok == token.DEFINE {
					add(c.Key, c.Value)
				}
			case *ast.BlockStmt:
				for _, statement := range c.List {
					for _, name := range declaredInStatement(statement) {
						result[name] = true
					}
				}
			}

			return true
		},
	)

	return result
}
