package multi_value_result

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
)

func calleeName(call *ast.CallExpr) string {
	switch fun := call.Fun.(type) {
	case *ast.Ident:
		return fun.Name
	case *ast.SelectorExpr:
		qualifier, okay := fun.X.(*ast.Ident)

		if okay {
			return join.Empty(
				qualifier.Name,
				constant.MemberSeparator,
				fun.Sel.Name,
			)
		}

		return fun.Sel.Name
	}

	return "a call"
}
