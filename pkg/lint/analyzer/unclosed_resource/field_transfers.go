package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"go/types"
)

func fieldTransfers(
	holder *types.Named,
	field *types.Var,
	value ast.Expr,
) []resource_transfer.Transfer {
	l, isLiteral := ast.Unparen(value).(*ast.CompositeLit)

	if !isLiteral || elementType(field.Type()) == nil {
		return []resource_transfer.Transfer{
			{Holder: holder, Field: field, Value: value},
		}
	}

	var result []resource_transfer.Transfer

	for _, element := range elementValues(l.Elts) {
		result = append(
			result,
			resource_transfer.Transfer{
				Holder:    holder,
				Field:     field,
				Container: true,
				Value:     element,
			},
		)
	}

	return result
}
