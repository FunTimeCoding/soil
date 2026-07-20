package assert_call

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

// A local test helper counts as an assert call when its name
// carries the assert prefix and it takes *testing.T first -
// its arguments are expected values one hop removed.
func isAssertHelperCall(
	p *packages.Package,
	call *ast.CallExpr,
) bool {
	i, okay := call.Fun.(*ast.Ident)

	if !okay || !strings.HasPrefix(i.Name, HelperPrefix) {
		return false
	}

	o := p.TypesInfo.ObjectOf(i)

	if o == nil {
		return false
	}

	f, okay := o.(*types.Func)

	if !okay {
		return false
	}

	s, okay := f.Type().(*types.Signature)

	if !okay || s.Params().Len() == 0 {
		return false
	}

	return s.Params().At(0).Type().String() == TestingType
}
