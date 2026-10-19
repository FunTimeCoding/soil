package error_answer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isErrorsFunction(
	p *packages.Package,
	name *ast.Ident,
	function string,
) bool {
	f, okay := p.TypesInfo.Uses[name].(*types.Func)

	return okay && f.Pkg() != nil &&
		f.Pkg().Path() == constant.ErrorsPackage && f.Name() == function
}
