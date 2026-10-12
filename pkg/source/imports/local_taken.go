package imports

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func LocalTaken(
	p *packages.Package,
	file *ast.File,
	name string,
) bool {
	for ident, o := range p.TypesInfo.Defs {
		if o == nil || ident.Name != name {
			continue
		}

		if !positionInFile(p.Fset, file, ident.Pos()) {
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

	return false
}
