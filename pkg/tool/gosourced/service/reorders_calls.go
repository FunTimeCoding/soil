package service

import (
	"go/ast"
	"slices"
)

func reordersCalls(
	values map[string]ast.Expr,
	source []string,
	planned []string,
) bool {
	calling := func(order []string) []string {
		var result []string

		for _, name := range order {
			if hasCall(values[name]) {
				result = append(result, name)
			}
		}

		return result
	}
	before := calling(source)

	return len(before) > 1 && !slices.Equal(before, calling(planned))
}
