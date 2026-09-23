package assert_call

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func IsAssertCall(
	p *packages.Package,
	call *ast.CallExpr,
) bool {
	s, okay := call.Fun.(*ast.SelectorExpr)

	if !okay {
		return isAssertHelperCall(p, call)
	}

	i, okay := s.X.(*ast.Ident)

	if !okay {
		return isAssertHelperCall(p, call)
	}

	n, okay := p.TypesInfo.ObjectOf(i).(*types.PkgName)

	if okay && n.Imported().Name() == constant.AssertPackageName {
		return true
	}

	return isAssertHelperCall(p, call)
}
