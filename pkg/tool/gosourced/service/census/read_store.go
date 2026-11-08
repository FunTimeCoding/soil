package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"slices"
)

func readStore(
	p *packages.Package,
	d *ast.FuncDecl,
	module string,
) *storeReading {
	result := newStoreReading()
	ast.Inspect(
		d.Body,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			f := called(p, c)

			if f == nil || owner(f) == nil {
				return true
			}

			o := owner(f)

			if o.Obj().Pkg().Path() == constant.GormPath &&
				o.Obj().Name() == constant.GormType {
				result.read = result.read ||
					slices.Contains(constant.GormReadVerbs, f.Name())
				result.written = result.written ||
					slices.Contains(constant.GormWriteVerbs, f.Name())

				if slices.Contains(constant.GormDestinationVerbs, f.Name()) {
					return true
				}

				for _, a := range c.Args {
					if m := model(p.TypesInfo.TypeOf(a), module); m != nil {
						result.models[qualified(m)] = m
					}
				}

				return true
			}

			if inModule(module, o.Obj().Pkg().Path()) {
				result.calls[methodKey(o, f.Name())] = true
			}

			return true
		},
	)

	return result
}
