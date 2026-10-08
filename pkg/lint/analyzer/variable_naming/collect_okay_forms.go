package variable_naming

import (
	"go/ast"
	"go/token"
	"go/types"
)

func collectOkayForms(
	y *types.Info,
	f *ast.FuncDecl,
) []OkayForm {
	var result []OkayForm

	if f.Body == nil {
		return result
	}

	process := func(s *ast.AssignStmt) {
		if s.Tok != token.DEFINE || len(s.Lhs) != 2 || len(s.Rhs) != 1 {
			return
		}

		if !isCommaOkayRHS(y, s.Rhs[0]) {
			return
		}

		ident, okay := s.Lhs[1].(*ast.Ident)

		if !okay || ident.Name == "_" {
			return
		}

		o := y.ObjectOf(ident)

		if o == nil {
			return
		}

		result = append(
			result,
			OkayForm{ident: ident, scope: o.Parent(), position: ident.Pos()},
		)
	}
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			if s, okay := n.(*ast.AssignStmt); okay {
				process(s)
			}

			return true
		},
	)

	return result
}
