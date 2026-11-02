package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func localFactoryRelays(
	p *packages.Package,
	f *ast.FuncDecl,
	indices map[*types.Var]int,
) []fact.Forward {
	var result []fact.Forward
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			a, okay := n.(*ast.AssignStmt)

			if !okay || len(a.Lhs) != len(a.Rhs) {
				return true
			}

			for i, l := range a.Lhs {
				s, okay := l.(*ast.SelectorExpr)

				if !okay {
					continue
				}

				local := rootIdentifier(p, s.X)

				if local == nil {
					continue
				}

				for _, c := range factoryCarriers(p, f.Body, a.Rhs[i], indices) {
					if c.Argument == constant.ReceiverArgument {
						continue
					}

					for _, sink := range passedOn(p, f.Body, local) {
						result = append(
							result,
							fact.Forward{
								Argument: c.Argument,
								Conditions: append(
									append([]fact.Condition{}, c.Conditions...),
									*fact.NewFieldDrawCondition(
										sink.Callee,
										sink.Argument,
										s.Sel.Name,
									),
								),
							},
						)
					}
				}
			}

			return true
		},
	)

	return result
}
