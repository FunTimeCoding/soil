package resource_transfer

import (
	"go/ast"
	"go/types"
)

type Transfer struct {
	Holder    *types.Named
	Field     *types.Var
	Container bool
	Value     ast.Expr
}
