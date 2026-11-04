package addition

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func NewTarget(
	a *Parameter,
	function *types.Func,
	owner *packages.Package,
	file *ast.File,
	declaration *ast.FuncDecl,
	t types.Type,
	qualifiers map[string]string,
) *Target {
	return &Target{
		Addition:    a,
		Function:    function,
		Owner:       owner,
		File:        file,
		Declaration: declaration,
		Type:        t,
		Qualifiers:  qualifiers,
	}
}
