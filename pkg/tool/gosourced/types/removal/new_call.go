package removal

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func NewCall(
	target *Target,
	owner *packages.Package,
	file *ast.File,
	call *ast.CallExpr,
	arguments []ast.Expr,
) *Call {
	return &Call{
		Target:    target,
		Owner:     owner,
		File:      file,
		Call:      call,
		Arguments: arguments,
	}
}
