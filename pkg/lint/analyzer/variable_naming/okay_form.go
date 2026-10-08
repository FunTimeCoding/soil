package variable_naming

import (
	"go/ast"
	"go/token"
	"go/types"
)

type OkayForm struct {
	ident    *ast.Ident
	scope    *types.Scope
	position token.Pos
}
