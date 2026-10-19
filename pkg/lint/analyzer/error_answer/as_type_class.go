package error_answer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func asTypeClass(
	p *packages.Package,
	x *ast.IfStmt,
) string {
	condition, okay := x.Cond.(*ast.Ident)

	if !okay {
		return ""
	}

	assignment, okay := x.Init.(*ast.AssignStmt)

	if !okay || len(assignment.Lhs) != 2 || len(assignment.Rhs) != 1 {
		return ""
	}

	flag, okay := assignment.Lhs[1].(*ast.Ident)

	if !okay || flag.Name != condition.Name {
		return ""
	}

	call, okay := assignment.Rhs[0].(*ast.CallExpr)

	if !okay {
		return ""
	}

	index, okay := call.Fun.(*ast.IndexExpr)

	if !okay {
		return ""
	}

	selector, okay := index.X.(*ast.SelectorExpr)

	if !okay || !isErrorsFunction(p, selector.Sel, constant.AsTypeFunction) {
		return ""
	}

	return classOfType(p.TypesInfo.TypeOf(index.Index))
}
