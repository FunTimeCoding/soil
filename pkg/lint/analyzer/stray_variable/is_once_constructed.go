package stray_variable

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isOnceConstructed(
	p *packages.Package,
	v *ast.ValueSpec,
	index int,
) bool {
	if index >= len(v.Values) {
		return false
	}

	call, okay := v.Values[index].(*ast.CallExpr)

	if !okay {
		return false
	}

	selector, okay := call.Fun.(*ast.SelectorExpr)

	if !okay {
		return false
	}

	x, okay := selector.X.(*ast.Ident)

	if !okay {
		return false
	}

	name, okay := p.TypesInfo.Uses[x].(*types.PkgName)

	return okay &&
		name.Imported().Path() == "sync" &&
		constant.OnceConstructors[selector.Sel.Name]
}
