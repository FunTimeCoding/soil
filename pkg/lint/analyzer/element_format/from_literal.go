package element_format

import "go/ast"

func FromLiteral(lit *ast.CompositeLit) *Elements {
	return &Elements{
		Open:     lit.Lbrace,
		Close:    lit.Rbrace,
		Items:    lit.Elts,
		Position: lit.Pos(),
	}
}
