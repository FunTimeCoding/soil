package unguarded_peer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	clients := guardedClients(p)

	if len(clients) == 0 {
		return
	}

	wrappers := guardingFunctions(p)
	inside := insideFunctions(p, wrappers)

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		var stack []ast.Node
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				if n == nil {
					stack = stack[:len(stack)-1]

					return false
				}

				stack = append(stack, n)
				s, okay := n.(*ast.SelectorExpr)

				if !okay {
					return true
				}

				selection := p.TypesInfo.Selections[s]

				if selection == nil ||
					selection.Kind() == types.FieldVal ||
					s.Sel.Name == constant.ReachableMethod ||
					s.Sel.Name == constant.CloseMethod ||
					!clients[holder(p, s.X)] ||
					!answersError(selection.Obj()) ||
					inGuard(p, stack, wrappers) ||
					inside[enclosing(p, stack)] {
					return true
				}

				results.AddConcern(
					concern.NewPosition(
						"unguarded_peer",
						fmt.Sprintf(
							"%s reaches a guarded peer outside its guard - call it inside the guard's Attempt or Do, peer.Guarded, or a function wrapping them",
							types.ExprString(s),
						),
						p.Fset.Position(s.Pos()),
					),
				)

				return true
			},
		)
	}
}
