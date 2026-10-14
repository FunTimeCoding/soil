package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func extractDisposals(
	p *packages.Package,
	f *ast.FuncDecl,
	o *types.Func,
	summary *fact.Summary,
) {
	indices := parameterIndices(o)

	if len(indices) == 0 {
		return
	}

	k := functionKey(o)
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			for parameter, index := range indices {
				refers := func(e ast.Expr) bool {
					return parameterOf(p, e, indices) == parameter
				}

				for _, condition := range callConditions(p, c, refers) {
					if condition.Callee == "" {
						summary.Disposes[k] = append(summary.Disposes[k], index)

						continue
					}

					summary.Forwards[k] = append(
						summary.Forwards[k],
						fact.Forward{
							Argument:   index,
							Conditions: []fact.Condition{condition},
						},
					)
				}
			}

			return true
		},
	)
}
