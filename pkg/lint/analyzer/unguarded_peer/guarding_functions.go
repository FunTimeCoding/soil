package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func guardingFunctions(p *packages.Package) map[types.Object]bool {
	result := make(map[types.Object]bool)

	for _, file := range p.Syntax {
		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Body == nil {
				continue
			}

			o := p.TypesInfo.Defs[f.Name]

			if o == nil ||
				!takesFunction(o.Type().(*types.Signature)) ||
				!callsGuard(p, f.Body) {
				continue
			}

			result[o] = true
		}
	}

	return result
}
