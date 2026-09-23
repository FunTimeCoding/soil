package unclosed_resource

import (
	"go/ast"
	"go/token"
	"go/types"
)

type candidate struct {
	object   *types.Var
	position token.Pos
	call     *ast.CallExpr
}
