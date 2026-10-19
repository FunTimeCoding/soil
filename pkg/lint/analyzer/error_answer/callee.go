package error_answer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func callee(
	p *packages.Package,
	call *ast.CallExpr,
) *types.Func {
	var name *ast.Ident

	if x, okay := call.Fun.(*ast.SelectorExpr); okay {
		name = x.Sel
	}

	if x, okay := call.Fun.(*ast.Ident); okay {
		name = x
	}

	if name == nil {
		return nil
	}

	f, _ := p.TypesInfo.Uses[name].(*types.Func)

	return f
}
