package symbol

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
)

func receiverIdentifiers(
	entries []*relocation.Entry,
) (map[*ast.Ident]bool, string) {
	result := map[*ast.Ident]bool{}

	for _, entry := range entries {
		declaration := entry.Declaration.(*ast.FuncDecl)
		expression := declaration.Recv.List[0].Type

		if star, okay := expression.(*ast.StarExpr); okay {
			expression = star.X
		}

		i, okay := expression.(*ast.Ident)

		if !okay {
			return nil, join.Empty(
				entry.Symbol,
				" has a generic receiver - move it by hand",
			)
		}

		result[i] = true
	}

	return result, ""
}
