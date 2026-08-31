package multi_value_result

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func calleeFunction(p *packages.Package, call *ast.CallExpr) *types.Func {
	var identifier *ast.Ident

	switch fun := call.Fun.(type) {
	case *ast.Ident:
		identifier = fun
	case *ast.SelectorExpr:
		identifier = fun.Sel
	default:
		return nil
	}

	f, okay := p.TypesInfo.Uses[identifier].(*types.Func)

	if !okay {
		return nil
	}

	return f
}
