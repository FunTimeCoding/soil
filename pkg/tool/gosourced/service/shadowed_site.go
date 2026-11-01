package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/token"
	"go/types"
)

func shadowedSite(
	q *relocation.FileQualification,
	targetPackagePath string,
) token.Pos {
	result := token.NoPos

	for ident := range q.Idents {
		scope := q.Owner.Types.Scope().Innermost(ident.Pos())

		if scope == nil {
			continue
		}

		_, o := scope.LookupParent(q.Name.Local, ident.Pos())

		if n, okay := o.(*types.PkgName); okay &&
			n.Imported().Path() == targetPackagePath {
			continue
		}

		if result == token.NoPos || ident.Pos() < result {
			result = ident.Pos()
		}
	}

	return result
}
