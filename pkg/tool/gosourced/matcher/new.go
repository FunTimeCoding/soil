package matcher

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/match"
	"go/ast"
	"go/token"
	"go/types"
)

func New(
	pattern *match.Pattern,
	information *types.Info,
	scope *types.Package,
	set *token.FileSet,
	querySymbol string,
	isAnchor func(types.Object) bool,
) *Matcher {
	return &Matcher{
		holes:       pattern.Holes,
		bindings:    map[string]ast.Expr{},
		information: information,
		scope:       scope,
		set:         set,
		querySymbol: querySymbol,
		isAnchor:    isAnchor,
	}
}
