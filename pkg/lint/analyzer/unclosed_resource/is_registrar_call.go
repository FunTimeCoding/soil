package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
)

func isRegistrarCall(c *ast.CallExpr) ast.Expr {
	selector, okay := c.Fun.(*ast.SelectorExpr)

	if !okay || !constant.CloseRegistrars[selector.Sel.Name] {
		return nil
	}

	if len(c.Args) != 1 {
		return nil
	}

	switch a := c.Args[0].(type) {
	case *ast.SelectorExpr:
		if a.Sel.Name == "Close" {
			return a.X
		}
	case *ast.FuncLit:
		return closureCloseSubject(a)
	}

	return nil
}
