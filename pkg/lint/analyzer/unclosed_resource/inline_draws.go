package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func inlineDraws(
	p *packages.Package,
	f *ast.FuncDecl,
	parameter *types.Var,
	index int,
) []fact.FieldForward {
	var result []fact.FieldForward

	for _, t := range returnedTransfers(p, f) {
		c, okay := ast.Unparen(t.Value).(*ast.CallExpr)

		if !okay {
			continue
		}

		field, okay := factoryField(p, c, parameter)

		if !okay {
			continue
		}

		result = append(
			result,
			fact.FieldForward{
				Argument: index,
				Field:    field,
				Conditions: []fact.Condition{
					*fact.NewReleaseCondition(
						typeKey(t.Holder.Origin().Obj()),
						t.Field.Name(),
					),
				},
			},
		)
	}

	return result
}
