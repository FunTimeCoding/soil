package expected_first

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func takesTestingFirst(
	p *packages.Package,
	f *ast.FuncDecl,
) bool {
	o := p.TypesInfo.Defs[f.Name]

	if o == nil {
		return false
	}

	s, okay := o.Type().(*types.Signature)

	if !okay || s.Params().Len() == 0 {
		return false
	}

	return s.Params().At(0).Type().String() == constant.AssertTestingType
}
