package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func readUse(
	p *packages.Package,
	i *ast.Ident,
	subject *types.Named,
	module string,
	r *reading.Reading,
) {
	o := p.TypesInfo.Uses[i]

	if o == nil || o.Pkg() == nil {
		return
	}

	path := o.Pkg().Path()

	if path != p.PkgPath && inModule(module, path) {
		r.Packages[relative(module, path)] = true
	}

	f, okay := o.(*types.Func)

	if !okay {
		return
	}

	n := owner(f)

	switch {
	case n != nil && sameType(n, subject):
		r.Calls[f.Name()] = true
	case n != nil && inModule(module, n.Obj().Pkg().Path()):
		l := label(n)

		if r.Uses[l] == nil {
			r.Uses[l] = map[string]bool{}
		}

		r.Uses[l][f.Name()] = true
		r.Keys[methodKey(n, f.Name())] = true
	case n == nil && path == p.PkgPath && f.Parent() == f.Pkg().Scope():
		r.Functions[f.Name()] = true
	}

	if n != nil && f.Name() == constant.AtomicallyMethod {
		r.Atomically = true
	}
}
