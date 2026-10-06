package removal

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func NewLocal(
	variable *types.Var,
	owner *packages.Package,
	statement ast.Stmt,
	block *ast.BlockStmt,
) *Local {
	return &Local{
		Variable:  variable,
		Owner:     owner,
		Statement: statement,
		Block:     block,
	}
}
