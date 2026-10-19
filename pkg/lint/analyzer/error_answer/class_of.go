package error_answer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func classOf(
	p *packages.Package,
	call *ast.CallExpr,
) string {
	if selector, okay := call.Fun.(*ast.SelectorExpr); okay && len(call.Args) == 2 &&
		isErrorsFunction(p, selector.Sel, constant.AsFunction) {
		return classOfType(p.TypesInfo.TypeOf(call.Args[1]))
	}

	f := callee(p, call)

	if f == nil || f.Pkg() == nil || f.Name() != constant.ErrorClassPredicate {
		return ""
	}

	return knownClass(f.Pkg().Path())
}
