package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func extractDraws(
	p *packages.Package,
	f *ast.FuncDecl,
	o *types.Func,
	summary *fact.Summary,
) {
	indices := parameterIndices(o)
	k := functionKey(o)

	for parameter, index := range indices {
		if index == constant.ReceiverArgument {
			continue
		}

		for drawn, field := range drawnLocals(p, f.Body, parameter) {
			for _, c := range drawnTakes(p, f, drawn) {
				summary.Draws[k] = append(
					summary.Draws[k],
					fact.FieldForward{
						Argument:   index,
						Field:      field,
						Conditions: c,
					},
				)
			}
		}

		summary.Draws[k] = append(
			summary.Draws[k],
			inlineDraws(p, f, parameter, index)...,
		)

		for _, h := range passedOn(p, f.Body, parameter) {
			summary.Draws[k] = append(
				summary.Draws[k],
				fact.FieldForward{
					Argument:   index,
					Field:      constant.AnyField,
					Conditions: []fact.Condition{*h},
				},
			)
		}
	}
}
