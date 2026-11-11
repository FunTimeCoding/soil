package symbol

import (
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
	"go/types"
)

func targetDeclares(
	plan *relocation.Plan,
	existing *ast.File,
	entries []*relocation.Entry,
	name string,
) bool {
	if existing != nil && imports.LocalTaken(plan.Target, existing, name) {
		return true
	}

	if plan.Target != nil && plan.Target.Types.Scope().Lookup(name) != nil {
		return true
	}

	for _, entry := range entries {
		if entry.NewName == name {
			return true
		}

		for ident, o := range plan.Source.TypesInfo.Defs {
			if o == nil || ident.Name != name ||
				ident.Pos() < entry.Declaration.Pos() ||
				ident.Pos() >= entry.Declaration.End() {
				continue
			}

			if _, okay := o.(*types.PkgName); okay {
				continue
			}

			if v, okay := o.(*types.Var); okay && v.IsField() {
				continue
			}

			return true
		}
	}

	return false
}
