package unclosed_resource

import "go/ast"

func elementValues(elements []ast.Expr) []ast.Expr {
	result := make([]ast.Expr, 0, len(elements))

	for _, e := range elements {
		if pair, okay := e.(*ast.KeyValueExpr); okay {
			result = append(result, pair.Value)

			continue
		}

		result = append(result, e)
	}

	return result
}
