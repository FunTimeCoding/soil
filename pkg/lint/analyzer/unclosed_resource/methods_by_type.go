package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func methodsByType(
	p *packages.Package,
) map[*types.TypeName]map[string]*ast.FuncDecl {
	result := make(map[*types.TypeName]map[string]*ast.FuncDecl)

	for _, file := range p.Syntax {
		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Recv == nil || f.Body == nil {
				continue
			}

			o, okay := p.TypesInfo.Defs[f.Name].(*types.Func)

			if !okay {
				continue
			}

			named, okay := dereference(o.Signature().Recv().Type()).(*types.Named)

			if !okay {
				continue
			}

			t := named.Origin().Obj()

			if result[t] == nil {
				result[t] = make(map[string]*ast.FuncDecl)
			}

			result[t][f.Name.Name] = f
		}
	}

	return result
}
