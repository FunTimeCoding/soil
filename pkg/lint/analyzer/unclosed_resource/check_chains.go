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

func checkChains(
	p *packages.Package,
	results *output.Results,
	body *ast.BlockStmt,
	s *Summaries,
) {
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			selector, okay := ast.Unparen(c.Fun).(*ast.SelectorExpr)

			if !okay {
				return true
			}

			owner, okay := ast.Unparen(selector.X).(*ast.CallExpr)

			if !okay || !carriesObligation(p.TypesInfo.TypeOf(owner)) ||
				!s.freshRoot(p, owner) ||
				s.isUnowned(
					p,
					body,
					resource_candidate.Candidate{Call: owner},
				) {
				return true
			}

			if f := calleeOf(p, c); f == nil || s.selfFor(f) || s.lendsFor(f) ||
				s.isReleasing(p, selector) {
				return true
			}

			results.AddConcern(
				concern.NewPosition(
					"unclosed_resource",
					fmt.Sprintf(
						"%s() is never closed; assign it and %s",
						types.ExprString(owner.Fun),
						chainSuggestion(p.TypesInfo.TypeOf(owner)),
					),
					p.Fset.Position(owner.Pos()),
				),
			)

			return true
		},
	)
}
