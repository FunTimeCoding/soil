package unclosed_resource

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func sharedVariables(p *packages.Package) []string {
	var result []string

	for _, file := range p.Syntax {
		for _, d := range file.Decls {
			g, okay := d.(*ast.GenDecl)

			if !okay || g.Tok != token.VAR {
				continue
			}

			for _, spec := range g.Specs {
				v, isValue := spec.(*ast.ValueSpec)

				if !isValue || len(v.Values) != len(v.Names) {
					continue
				}

				for i, name := range v.Names {
					o, isVariable := p.TypesInfo.Defs[name].(*types.Var)

					if isVariable && isSharedConstructor(p, v.Values[i]) {
						result = append(result, variableKey(o))
					}
				}
			}
		}
	}

	return result
}
