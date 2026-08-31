package match

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Matcher struct {
	holes       map[string]string
	bindings    map[string]ast.Expr
	information *types.Info
	scope       *types.Package
	set         *token.FileSet
	querySymbol string
	isAnchor    func(types.Object) bool
}
