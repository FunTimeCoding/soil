package assert_call

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
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
		return false
	}

	o := p.TypesInfo.ObjectOf(i)

	if o == nil {
		return false
	}

	n, okay := o.(*types.PkgName)

	if !okay {
		return false
	}

	return strings.HasSuffix(n.Imported().Path(), PackageSuffix)
}
