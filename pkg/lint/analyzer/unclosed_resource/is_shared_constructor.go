package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func isSharedConstructor(
	p *packages.Package,
	e ast.Expr,
) bool {
	c, okay := ast.Unparen(e).(*ast.CallExpr)

	if !okay {
		return false
	}

	f := calleeOf(p, c)

	return f != nil && f.Pkg() != nil &&
		f.Pkg().Path() == constant.SyncPackage &&
		constant.SharedConstructors[f.Name()]
}
