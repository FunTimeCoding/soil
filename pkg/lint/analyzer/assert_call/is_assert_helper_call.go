package assert_call

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func isAssertHelperCall(
	p *packages.Package,
	call *ast.CallExpr,
) bool {
	f := calledFunction(p, call)

	if f == nil || !HasAssertPrefix(f.Name()) {
		return false
	}

	return takesTestingFirst(f) || isMethod(f)
}
