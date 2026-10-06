package service

import (
	"go/ast"
	"go/types"
)

func literalFields(
	structure *types.Struct,
	lit *ast.CompositeLit,
) ([]string, bool) {
	set := map[string]bool{}
	unkeyed := false

	for i, element := range lit.Elts {
		pair, okay := element.(*ast.KeyValueExpr)

		if !okay {
			unkeyed = true

			if i < structure.NumFields() {
				set[structure.Field(i).Name()] = true
			}

			continue
		}

		if key, okay := pair.Key.(*ast.Ident); okay {
			set[key.Name] = true
		}
	}

	var result []string

	for i := range structure.NumFields() {
		if name := structure.Field(i).Name(); set[name] {
			result = append(result, name)
		}
	}

	return result, unkeyed
}
