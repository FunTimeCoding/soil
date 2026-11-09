package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/call_sink"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func callSinks(
	p *packages.Package,
	body *ast.BlockStmt,
) map[*ast.CallExpr]call_sink.Sink {
	result := map[*ast.CallExpr]call_sink.Sink{}
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			g := callee(p, c)

			if g == nil {
				return true
			}

			for j, a := range c.Args {
				sink := call_sink.Sink{Callee: g, Argument: j}

				switch v := ast.Unparen(a).(type) {
				case *ast.CallExpr:
					result[v] = sink
				case *ast.Ident:
					if o, okay := p.TypesInfo.Uses[v].(*types.Var); okay {
						for _, e := range assignedValues(p, body, o) {
							if s, okay := ast.Unparen(e).(*ast.CallExpr); okay {
								result[s] = sink
							}
						}
					}
				}
			}

			return true
		},
	)

	return result
}
