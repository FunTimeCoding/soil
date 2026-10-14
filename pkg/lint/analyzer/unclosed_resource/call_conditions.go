package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func callConditions(
	p *packages.Package,
	c *ast.CallExpr,
	refers func(ast.Expr) bool,
) []fact.Condition {
	var result []fact.Condition

	if selector, okay := ast.Unparen(c.Fun).(*ast.SelectorExpr); okay {
		f, isFunction := p.TypesInfo.Uses[selector.Sel].(*types.Func)

		if isFunction && f.Signature().Recv() != nil && refers(selector.X) {
			result = append(result, *methodCondition(f))
		}
	}

	f := calleeOf(p, c)

	if f == nil {
		return result
	}

	for i, a := range c.Args {
		if refers(a) {
			result = append(result, *argumentCondition(f, i))
		}
	}

	return result
}
