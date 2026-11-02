package removal

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Target struct {
	Function    *types.Func
	Owner       *packages.Package
	File        *ast.File
	Declaration *ast.FuncDecl
	Indices     map[int]bool
	Variables   []*types.Var
	Variadic    int
}
