package variable_naming

import (
	"go/ast"
	"go/types"
)

func collectFromRange(
	y *types.Info,
	s *ast.RangeStmt,
	result *[]TypedVariable,
) {
	for _, expr := range []ast.Expr{s.Key, s.Value} {
		if expr == nil {
			continue
		}

		ident, okay := expr.(*ast.Ident)

		if !okay || ident.Name == "_" {
			continue
		}

		o := y.ObjectOf(ident)

		if o == nil {
			continue
		}

		*result = append(
			*result,
			TypedVariable{
				ident:       ident,
				typ:         o.Type(),
				precedence:  typePrecedence(o.Type()),
				scopedNames: collectScopedNames(o),
			},
		)
	}
}
