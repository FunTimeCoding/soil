package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func drawnTakes(
	p *packages.Package,
	f *ast.FuncDecl,
	drawn *types.Var,
) [][]fact.Condition {
	indices := map[*types.Var]int{drawn: 0}
	var result [][]fact.Condition

	for _, t := range returnedTransfers(p, f) {
		for _, c := range carriers(
			p,
			f.Body,
			t.Value,
			indices,
			map[*types.Var]bool{},
		) {
			result = append(
				result,
				append(
					[]fact.Condition{
						*fact.NewReleaseCondition(
							typeKey(t.Holder.Origin().Obj()),
							t.Field.Name(),
						),
					},
					c.Conditions...,
				),
			)
		}
	}

	refers := func(e ast.Expr) bool {
		return parameter(p, e, indices) == drawn
	}
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.CallExpr:
				for _, condition := range callConditions(p, v, refers) {
					result = append(result, []fact.Condition{condition})
				}
			case *ast.SelectorExpr:
				f, _ := p.TypesInfo.Uses[v.Sel].(*types.Func)

				if isCloseMethod(f) && refers(v.X) {
					result = append(result, nil)
				}
			}

			return true
		},
	)

	return result
}
