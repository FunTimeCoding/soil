package symbol

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/types"
	"sort"
)

func shadowingLocals(
	q *relocation.FileQualification,
	name string,
) ([]*types.Var, bool) {
	if q.Owner.Types.Scope().Lookup(name) != nil {
		return nil, false
	}

	var result []*types.Var

	for ident, o := range q.Owner.TypesInfo.Defs {
		if o == nil || ident.Name != name ||
			ident.Pos() < q.File.Pos() || ident.Pos() >= q.File.End() {
			continue
		}

		if _, okay := o.(*types.PkgName); okay {
			continue
		}

		v, okay := o.(*types.Var)

		if !okay || v.IsField() || v.Parent() == nil ||
			v.Parent() == q.Owner.Types.Scope() {
			return nil, false
		}

		result = append(result, v)
	}

	sort.Slice(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Pos() < result[j].Pos()
		},
	)

	return result, len(result) > 0
}
