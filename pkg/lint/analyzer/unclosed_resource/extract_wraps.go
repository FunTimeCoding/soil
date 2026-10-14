package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func extractWraps(
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

	for _, t := range returnedTransfers(p, f) {
		o2 := t.Holder.Origin().Obj()

		for _, c := range carriers(
			p,
			f.Body,
			t.Value,
			indices,
			map[*types.Var]bool{},
		) {
			if c.Argument == constant.ReceiverArgument {
				continue
			}

			summary.Wraps[k] = append(
				summary.Wraps[k],
				fact.Forward{
					Argument: c.Argument,
					Conditions: append(
						[]fact.Condition{
							*fact.NewReleaseCondition(
								typeKey(o2),
								t.Field.Name(),
							),
						},
						c.Conditions...,
					),
				},
			)
		}
	}
}
