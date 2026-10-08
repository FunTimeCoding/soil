package variable_naming

import (
	"go/ast"
	"go/token"
	"go/types"
)

func collectFromAssign(
	y *types.Info,
	s *ast.AssignStmt,
	result *[]TypedVariable,
) {
	if s.Tok != token.DEFINE {
		return
	}

	commaOkay := len(s.Lhs) == 2 && len(s.Rhs) == 1 && isCommaOkayRHS(
		y,
		s.Rhs[0],
	)

	for i, lhs := range s.Lhs {
		ident, okay := lhs.(*ast.Ident)

		if !okay || ident.Name == "_" {
			continue
		}

		if commaOkay && i == 1 {
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
