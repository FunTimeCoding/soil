package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func passedOn(
	p *packages.Package,
	body *ast.BlockStmt,
	parameter *types.Var,
) []*fact.Condition {
	var result []*fact.Condition
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			h := callee(p, c)

			if h == nil {
				return true
			}

			for j, a := range c.Args {
				if i, okay := ast.Unparen(a).(*ast.Ident); okay &&
					p.TypesInfo.Uses[i] == parameter {
					result = append(
						result,
						fact.NewDrawCondition(functionKey(h), j),
					)
				}
			}

			return true
		},
	)

	return result
}
