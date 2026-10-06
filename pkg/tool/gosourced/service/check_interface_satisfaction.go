package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkInterfaceSatisfaction(
	all []*packages.Package,
	set *token.FileSet,
	t *removal.Target,
	r *output.Results,
) {
	receiver := t.Function.Type().(*types.Signature).Recv()

	if receiver == nil {
		return
	}

	base := receiver.Type()

	if p, okay := base.(*types.Pointer); okay {
		base = p.Elem()
	}

	pointer := types.NewPointer(base)
	seen := map[*types.TypeName]bool{}

	for _, p := range all {
		for _, name := range p.Types.Scope().Names() {
			o, okay := p.Types.Scope().Lookup(name).(*types.TypeName)

			if !okay {
				continue
			}

			i, isInterface := o.Type().Underlying().(*types.Interface)

			if !isInterface || seen[o] {
				continue
			}

			seen[o] = true
			declares := false

			for m := range i.Methods() {
				if m.Name() == t.Function.Name() {
					declares = true
				}
			}

			if !declares ||
				!types.Implements(base, i) && !types.Implements(pointer, i) {
				continue
			}

			addPositionConcern(
				r,
				set,
				t.Function.Pos(),
				constant.ConcernValueUse,
				false,
				"%s satisfies %s.%s; its signature cannot change",
				t.Function.Name(),
				p.PkgPath,
				name,
			)
		}
	}
}
