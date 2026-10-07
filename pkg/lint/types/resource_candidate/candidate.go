package resource_candidate

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Candidate struct {
	Object   *types.Var
	Position token.Pos
	Call     *ast.CallExpr
}
