package addition

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Target struct {
	Addition    *Parameter
	Function    *types.Func
	Owner       *packages.Package
	File        *ast.File
	Declaration *ast.FuncDecl
	Type        types.Type
	Qualifiers  map[string]string
}
