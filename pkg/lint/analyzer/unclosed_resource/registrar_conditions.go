package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func registrarConditions(
	p *packages.Package,
	body *ast.BlockStmt,
) []fact.Condition {
	flows := func(e ast.Expr) bool {
		o := rootIdentifier(p, e)

		return o != nil && flowsToResult(p, body, o)
	}
	var result []fact.Condition
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay || len(c.Args) != 1 {
				return true
			}

			selector, okay := c.Fun.(*ast.SelectorExpr)

			if !okay || !constant.CloseRegistrars[selector.Sel.Name] {
				return true
			}

			switch a := c.Args[0].(type) {
			case *ast.SelectorExpr:
				f, isFunction := p.TypesInfo.Uses[a.Sel].(*types.Func)

				if isFunction && f.Signature().Recv() != nil && flows(a.X) {
					result = append(result, *methodCondition(f))
				}
			case *ast.FuncLit:
				if call := singleCall(a); call != nil {
					result = append(result, callConditions(p, call, flows)...)
				}
			}

			return true
		},
	)

	return result
}
