package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/ast"
	"go/types"
)

func sourceStays(
	q *relocation.FileQualification,
	sourcePackagePath string,
) bool {
	result := false
	ast.Inspect(
		q.File,
		func(n ast.Node) bool {
			selector, okay := n.(*ast.SelectorExpr)

			if !okay || result {
				return !result
			}

			x, okay := selector.X.(*ast.Ident)

			if !okay {
				return true
			}

			name, okay := q.Owner.TypesInfo.Uses[x].(*types.PkgName)

			if !okay || name.Imported().Path() != sourcePackagePath {
				return true
			}

			if _, moving := q.Idents[selector.Sel]; !moving {
				result = true
			}

			return true
		},
	)

	return result
}
