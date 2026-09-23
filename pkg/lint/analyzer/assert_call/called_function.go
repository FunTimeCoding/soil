package assert_call

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func calledFunction(
	p *packages.Package,
	call *ast.CallExpr,
) *types.Func {
	var name *ast.Ident

	switch v := call.Fun.(type) {
	case *ast.Ident:
		name = v
	case *ast.SelectorExpr:
		name = v.Sel
	default:
		return nil
	}

	o := p.TypesInfo.ObjectOf(name)

	if o == nil {
		return nil
	}

	f, okay := o.(*types.Func)

	if !okay {
		return nil
	}

	return f
}
