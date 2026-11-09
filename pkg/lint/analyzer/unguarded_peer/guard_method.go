package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func guardMethod(
	p *packages.Package,
	f ast.Expr,
) bool {
	s, okay := f.(*ast.SelectorExpr)

	if !okay {
		return false
	}

	selection := p.TypesInfo.Selections[s]

	return selection != nil &&
		selection.Kind() == types.MethodVal &&
		isGuard(selection.Recv()) &&
		takesFunction(selection.Obj().Type().(*types.Signature))
}
