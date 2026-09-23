package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
)

func closeSubject(c *ast.CallExpr) ast.Expr {
	selector, okay := c.Fun.(*ast.SelectorExpr)

	if !okay {
		return nil
	}

	if selector.Sel.Name == "Close" && len(c.Args) == 0 {
		return selector.X
	}

	if constant.CloseHelpers[selector.Sel.Name] && len(c.Args) > 0 {
		return c.Args[0]
	}

	return nil
}
