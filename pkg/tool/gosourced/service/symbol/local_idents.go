package symbol

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
	"go/types"
)

func localIdents(
	q *relocation.FileQualification,
	v *types.Var,
) []*ast.Ident {
	var result []*ast.Ident

	for ident, o := range q.Owner.TypesInfo.Defs {
		if o == v {
			result = append(result, ident)
		}
	}

	for ident, o := range q.Owner.TypesInfo.Uses {
		if o == v {
			result = append(result, ident)
		}
	}

	return result
}
