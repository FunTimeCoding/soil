package string_concatenation

import (
	"go/ast"
	"go/token"
	"strings"
)

func containsRawString(e ast.Expr) bool {
	result := false
	ast.Inspect(
		e,
		func(n ast.Node) bool {
			l, okay := n.(*ast.BasicLit)

			if okay && l.Kind == token.STRING &&
				strings.HasPrefix(l.Value, "`") {
				result = true
			}

			return !result
		},
	)

	return result
}
