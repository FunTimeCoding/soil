package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func extractYields(
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
	returned := returnedVariables(p, f)
	yield := func(field string, e ast.Expr) {
		for _, c := range factoryCarriers(p, f.Body, e, indices) {
			if c.Argument != constant.ReceiverArgument {
				summary.Yields[k] = append(
					summary.Yields[k],
					fact.FieldForward{
						Argument:   c.Argument,
						Field:      field,
						Conditions: c.Conditions,
					},
				)
			}
		}
	}
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.ReturnStmt:
				for _, r := range v.Results {
					yield("", r)

					for field, e := range literalFields(r) {
						yield(field, e)
					}
				}
			case *ast.AssignStmt:
				for i, l := range v.Lhs {
					s, okay := l.(*ast.SelectorExpr)

					if okay && i < len(v.Rhs) && returnedRoot(p, s.X, returned) {
						yield(s.Sel.Name, v.Rhs[i])
					}
				}
			}

			return true
		},
	)
}
