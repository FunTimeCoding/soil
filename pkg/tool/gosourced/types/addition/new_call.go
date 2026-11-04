package addition

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func NewCall(
	t *Target,
	owner *packages.Package,
	file *ast.File,
	call *ast.CallExpr,
	qualifiers map[string]string,
) *Call {
	return &Call{
		Target:     t,
		Owner:      owner,
		File:       file,
		Call:       call,
		Qualifiers: qualifiers,
	}
}
