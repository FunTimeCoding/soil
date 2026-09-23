package assert_call

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func ExpectedIndex(
	p *packages.Package,
	call *ast.CallExpr,
) int {
	f := calledFunction(p, call)

	if f == nil {
		return 0
	}

	if takesTestingFirst(f) {
		return 1
	}

	return 0
}
