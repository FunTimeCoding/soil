package match

import (
	"fmt"
	"go/ast"
)

func CheckApplyShape(
	pattern *Pattern,
	replacement *Pattern,
	querySymbol string,
) error {
	for name := range replacement.Holes {
		if _, known := pattern.Holes[name]; !known {
			return fmt.Errorf(
				"replacement hole %s is not bound by the pattern",
				name,
			)
		}
	}

	for index, p := range []*Pattern{pattern, replacement} {
		blocks := 0
		anchors := 0
		ast.Inspect(
			p.Statement,
			func(n ast.Node) bool {
				if _, isBlock := n.(*ast.BlockStmt); isBlock {
					blocks++
				}

				if ident, isIdent := n.(*ast.Ident); isIdent &&
					ident.Name == querySymbol {
					anchors++
				}

				return true
			},
		)

		if blocks > 0 {
			return fmt.Errorf(
				"apply does not rewrite statements containing blocks yet",
			)
		}

		if index == 0 && anchors != 1 {
			return fmt.Errorf(
				"the pattern must reference the anchor %s exactly once",
				querySymbol,
			)
		}

		if index == 1 && anchors > 1 {
			return fmt.Errorf(
				"the replacement may reference the anchor %s at most once",
				querySymbol,
			)
		}
	}

	return nil
}
