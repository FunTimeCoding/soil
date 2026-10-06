package removal

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func NewTarget(
	function *types.Func,
	owner *packages.Package,
	file *ast.File,
	declaration *ast.FuncDecl,
) *Target {
	variadic := -1
	s := function.Type().(*types.Signature)

	if s.Variadic() {
		variadic = s.Params().Len() - 1
	}

	return &Target{
		Function:    function,
		Owner:       owner,
		File:        file,
		Declaration: declaration,
		Indices:     map[int]bool{},
		Variadic:    variadic,
	}
}
