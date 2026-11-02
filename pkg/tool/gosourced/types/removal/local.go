package removal

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Local struct {
	Variable  *types.Var
	Owner     *packages.Package
	Statement ast.Stmt
	Block     *ast.BlockStmt
}
