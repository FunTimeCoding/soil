package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Extract(p *packages.Package) *fact.Summary {
	result := fact.NewSummary()

	if p.TypesInfo == nil {
		return result
	}

	for _, file := range p.Syntax {
		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Body == nil {
				continue
			}

			o, okay := p.TypesInfo.Defs[f.Name].(*types.Func)

			if !okay {
				continue
			}

			k := functionKey(o)

			if directArranges(p, f.Body) {
				result.Arranges = append(result.Arranges, k)

				continue
			}

			for _, g := range delegatesOf(p, f.Body) {
				if callee := functionKey(g); callee != k {
					result.Delegates[k] = append(result.Delegates[k], callee)
				}
			}
		}
	}

	return result
}
