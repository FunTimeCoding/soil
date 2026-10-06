package service

import (
	"go/ast"
	"go/types"
)

func siteValues(
	structure *types.Struct,
	target ast.Expr,
) (map[string]ast.Expr, []string) {
	result := map[string]ast.Expr{}
	var order []string
	lit, okay := target.(*ast.CompositeLit)

	if !okay {
		return result, order
	}

	for i, element := range lit.Elts {
		name := ""
		value := element

		if pair, isPair := element.(*ast.KeyValueExpr); isPair {
			if key, isIdent := pair.Key.(*ast.Ident); isIdent {
				name = key.Name
			}

			value = pair.Value
		} else if i < structure.NumFields() {
			name = structure.Field(i).Name()
		}

		result[name] = value
		order = append(order, name)
	}

	return result, order
}
