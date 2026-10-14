package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func extractTakes(
	p *packages.Package,
	f *ast.FuncDecl,
	o *types.Func,
	summary *fact.Summary,
) {
	r := receiverOf(p, f)

	if r == nil {
		return
	}

	named, okay := dereference(r.Type()).(*types.Named)

	if !okay {
		return
	}

	holder := typeKey(named.Origin().Obj())
	indices := parameterIndices(o)
	k := functionKey(o)
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			a, isAssign := n.(*ast.AssignStmt)

			if !isAssign || len(a.Lhs) != len(a.Rhs) {
				return true
			}

			for i, l := range a.Lhs {
				target := l

				if index, isIndex := l.(*ast.IndexExpr); isIndex {
					target = index.X
				}

				field := fieldOfReceiver(p, target, r)

				if field == "" {
					continue
				}

				for _, v := range storedValues(p, target, a.Rhs[i]) {
					for _, c := range carriers(
						p,
						f.Body,
						v,
						indices,
						map[*types.Var]bool{},
					) {
						if c.Argument == constant.ReceiverArgument {
							continue
						}

						summary.Forwards[k] = append(
							summary.Forwards[k],
							fact.Forward{
								Argument: c.Argument,
								Conditions: append(
									[]fact.Condition{
										*fact.NewReleaseCondition(holder, field),
									},
									c.Conditions...,
								),
							},
						)
					}
				}
			}

			return true
		},
	)
}
