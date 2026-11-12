package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/types"
)

func leavingTaken(
	q *relocation.FileQualification,
	name string,
) (bool, bool) {
	if !q.SamePackage {
		return false, false
	}

	held := q.Owner.Types.Scope().Lookup(name)

	if held == nil {
		return false, false
	}

	leaving := false

	for ident := range q.Idents {
		if q.Owner.TypesInfo.Uses[ident] == held {
			leaving = true
		}
	}

	if !leaving {
		return false, false
	}

	for ident, o := range q.Owner.TypesInfo.Defs {
		if o == nil || o == held || ident.Name != name ||
			ident.Pos() < q.File.Pos() || ident.Pos() > q.File.End() {
			continue
		}

		if _, okay := o.(*types.PkgName); okay {
			continue
		}

		if v, okay := o.(*types.Var); okay && v.IsField() {
			continue
		}

		return true, true
	}

	return false, true
}
