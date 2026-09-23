package forwarding_function

import (
	"fmt"
	"go/ast"
)

func calledName(e ast.Expr) string {
	switch v := e.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.SelectorExpr:
		identifier, okay := v.X.(*ast.Ident)

		if !okay {
			return v.Sel.Name
		}

		return fmt.Sprintf("%s.%s", identifier.Name, v.Sel.Name)
	}

	return ""
}
