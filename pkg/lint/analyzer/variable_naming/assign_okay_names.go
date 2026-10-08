package variable_naming

import (
	"fmt"
	"go/ast"
	"go/types"
	"sort"
)

func assignOkayNames(forms []OkayForm) map[*ast.Ident]string {
	result := map[*ast.Ident]string{}
	byScope := map[*types.Scope][]OkayForm{}

	for _, f := range forms {
		byScope[f.scope] = append(byScope[f.scope], f)
	}

	for _, scopeForms := range byScope {
		sort.Slice(
			scopeForms,
			func(i, j int) bool {
				return scopeForms[i].position < scopeForms[j].position
			},
		)

		for i, f := range scopeForms {
			name := "okay"

			if i > 0 {
				name = fmt.Sprintf("okay%d", i)
			}

			result[f.ident] = name
		}
	}

	return result
}
