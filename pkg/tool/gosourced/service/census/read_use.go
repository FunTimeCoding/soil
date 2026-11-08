package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func readUse(
	p *packages.Package,
	i *ast.Ident,
	subject *types.Named,
	module string,
	r *reading,
) {
	o := p.TypesInfo.Uses[i]

	if o == nil || o.Pkg() == nil {
		return
	}

	path := o.Pkg().Path()

	if path != p.PkgPath && inModule(module, path) {
		r.packages[relative(module, path)] = true
	}

	f, okay := o.(*types.Func)

	if !okay {
		return
	}

	n := owner(f)

	switch {
	case n != nil && sameType(n, subject):
		r.calls[f.Name()] = true
	case n != nil && inModule(module, n.Obj().Pkg().Path()):
		l := label(n)

		if r.uses[l] == nil {
			r.uses[l] = map[string]bool{}
		}

		r.uses[l][f.Name()] = true
		r.keys[methodKey(n, f.Name())] = true
	case n == nil && path == p.PkgPath && f.Parent() == f.Pkg().Scope():
		r.functions[f.Name()] = true
	}

	if n != nil && f.Name() == constant.AtomicallyMethod {
		r.atomically = true
	}
}
