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
			if s, okay := n.(*ast.SelectorExpr); okay {
				summary.Disposes[k] = append(
					summary.Disposes[k],
					registeredCloses(p, s, indices)...,
				)
			}

			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			for v, index := range indices {
				refers := func(e ast.Expr) bool {
					return parameter(p, e, indices) == v
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
