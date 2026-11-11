package symbol

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkInterfaceSatisfaction(
	all []*packages.Package,
	set *token.FileSet,
	f *types.Func,
	r *output.Results,
) {
	receiver := f.Type().(*types.Signature).Recv()

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
				if m.Name() == f.Name() {
					declares = true
				}
			}

			if !declares ||
				!types.Implements(base, i) && !types.Implements(pointer, i) {
				continue
			}

			concern_line.Add(
				r,
				set,
				f.Pos(),
				constant.ConcernValueUse,
				false,
				"%s satisfies %s.%s; its signature cannot change",
				f.Name(),
				p.PkgPath,
				name,
			)
		}
	}
}
