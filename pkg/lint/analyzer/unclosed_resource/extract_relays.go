package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func extractRelays(
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
	summary.Forwards[k] = append(
		summary.Forwards[k],
		localFactoryRelays(p, f, indices)...,
	)

	for source, sink := range callSinks(p, f.Body) {
		from := callee(p, source)

		if from == nil {
			continue
		}

		for i, a := range source.Args {
			for _, c := range carriers(
				p,
				f.Body,
				a,
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
							append([]fact.Condition{}, c.Conditions...),
							*fact.NewRelayCondition(
								functionKey(from),
								i,
								functionKey(sink.Callee),
								sink.Argument,
							),
						),
					},
				)
			}
		}
	}
}
