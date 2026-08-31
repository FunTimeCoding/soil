package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/match"
	"go/ast"
)

func resolveQualifiers(
	replacement *match.Pattern,
	querySymbol string,
	names map[string][]string,
) (map[string]string, error) {
	result := map[string]string{}
	var failure error
	ast.Inspect(
		replacement.Statement,
		func(n ast.Node) bool {
			selector, okay := n.(*ast.SelectorExpr)

			if !okay || selector.Sel.Name == querySymbol {
				return true
			}

			qualifier, plain := selector.X.(*ast.Ident)

			if !plain {
				return true
			}

			if _, isHole := replacement.Holes[qualifier.Name]; isHole {
				return true
			}

			if path, given := replacement.Imports[qualifier.Name]; given {
				result[qualifier.Name] = path

				return true
			}

			candidates := names[qualifier.Name]

			if len(candidates) == 1 {
				result[qualifier.Name] = candidates[0]

				return true
			}

			if failure == nil {
				failure = fmt.Errorf(
					"cannot resolve package %s (%d candidates) - pass its import line in the pattern source",
					qualifier.Name,
					len(candidates),
				)
			}

			return true
		},
	)

	return result, failure
}
