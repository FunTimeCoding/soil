package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func callRelease(
	p *packages.Package,
	c *ast.CallExpr,
	r *types.Var,
) []fact.Release {
	var result []fact.Release

	if f := calleeOf(p, c); f != nil {
		for i, a := range c.Args {
			if field := fieldOfReceiver(p, a, r); field != "" {
				result = append(
					result,
					fact.Release{
						Field:     field,
						Condition: *argumentCondition(f, i),
					},
				)
			}
		}
	}

	selector, okay := c.Fun.(*ast.SelectorExpr)

	if !okay {
		return result
	}

	method, _ := p.TypesInfo.Uses[selector.Sel].(*types.Func)

	if method == nil || method.Signature().Recv() == nil {
		return result
	}

	condition := *methodCondition(method)

	if isReceiver(p, selector.X, r) {
		s := p.TypesInfo.Selections[selector]

		if s == nil || len(s.Index()) < 2 {
			return result
		}

		return append(
			result,
			fact.Release{
				Field:     fieldName(r.Type(), s.Index()[0]),
				Condition: condition,
			},
		)
	}

	field := fieldOfReceiver(p, selector.X, r)

	if field == "" {
		return result
	}

	inner, direct := selector.X.(*ast.SelectorExpr)

	if condition.Callee != "" && (!direct || !isReceiver(p, inner.X, r)) {
		return result
	}

	return append(result, fact.Release{Field: field, Condition: condition})
}
