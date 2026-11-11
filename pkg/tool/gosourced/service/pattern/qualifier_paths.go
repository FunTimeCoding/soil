package pattern

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/match"
	"go/ast"
)

func qualifierPaths(
	p *match.Pattern,
	querySymbol string,
	anchor string,
	names map[string][]string,
) (map[string]string, map[string]int) {
	result := map[string]string{}
	unresolved := map[string]int{}
	ast.Inspect(
		p.Statement,
		func(n ast.Node) bool {
			selector, okay := n.(*ast.SelectorExpr)

			if !okay {
				return true
			}

			qualifier, plain := selector.X.(*ast.Ident)

			if !plain {
				return true
			}

			if _, isHole := p.Holes[qualifier.Name]; isHole {
				return true
			}

			if selector.Sel.Name == querySymbol {
				if anchor != "" {
					result[qualifier.Name] = anchor
				}

				return true
			}

			if path, given := p.Imports[qualifier.Name]; given {
				result[qualifier.Name] = path

				return true
			}

			candidates := names[qualifier.Name]

			if len(candidates) == 1 {
				result[qualifier.Name] = candidates[0]

				return true
			}

			unresolved[qualifier.Name] = len(candidates)

			return true
		},
	)

	return result, unresolved
}
