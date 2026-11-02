package unclosed_resource

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/types/resource_candidate"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkArguments(
	p *packages.Package,
	results *output.Results,
	body *ast.BlockStmt,
	s *Summaries,
) {
	sinks := callSinks(p, body)
	returned := returnedCalls(p, body)
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			f := calleeOf(p, c)

			if f == nil {
				return true
			}

			for i, a := range c.Args {
				owner, okay := ast.Unparen(a).(*ast.CallExpr)

				if !okay || s.takes(f, i) ||
					s.handedOn(c, f, i, sinks, returned) ||
					!carriesObligation(p.TypesInfo.TypeOf(owner)) ||
					!s.freshRoot(p, owner) ||
					s.isUnowned(
						p,
						body,
						resource_candidate.Candidate{Call: owner},
					) {
					continue
				}

				results.AddConcern(
					concern.NewPosition(
						"unclosed_resource",
						fmt.Sprintf(
							"%s() passed to %s is never closed; assign it and %s",
							types.ExprString(owner.Fun),
							types.ExprString(c.Fun),
							chainSuggestion(p.TypesInfo.TypeOf(owner)),
						),
						p.Fset.Position(owner.Pos()),
					),
				)
			}

			return true
		},
	)
}
