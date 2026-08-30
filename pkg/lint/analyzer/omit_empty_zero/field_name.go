package omit_empty_zero

import (
	"go/ast"
	"go/types"
)

func fieldName(field *ast.Field) string {
	if len(field.Names) > 0 {
		return field.Names[0].Name
	}

	return types.ExprString(field.Type)
}
