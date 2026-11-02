package match

import (
	"fmt"
	"go/ast"
)

func checkSpreadHoles(
	statement ast.Stmt,
	holes map[string]string,
) error {
	var failure error
	ast.Inspect(
		statement,
		func(n ast.Node) bool {
			call, okay := n.(*ast.CallExpr)

			if !okay || len(call.Args) != 1 || !call.Ellipsis.IsValid() {
				return true
			}

			hole, wildcard := call.Args[0].(*ast.Ident)

			if !wildcard {
				return true
			}

			declared, isHole := holes[hole.Name]

			if isHole && declared != "[]any" && failure == nil {
				failure = fmt.Errorf(
					"spread hole %s must be declared []any - it matches any argument list and checks nothing",
					hole.Name,
				)
			}

			return true
		},
	)

	return failure
}
