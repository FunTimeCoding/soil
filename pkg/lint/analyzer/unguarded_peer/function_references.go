package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func functionReferences(
	p *packages.Package,
	wrappers map[types.Object]bool,
) map[types.Object][]*Reference {
	result := make(map[types.Object][]*Reference)

	for _, file := range p.Syntax {
		var stack []ast.Node
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				if n == nil {
					stack = stack[:len(stack)-1]

					return false
				}

				stack = append(stack, n)
				i, okay := n.(*ast.Ident)

				if !okay {
					return true
				}

				f, isFunction := p.TypesInfo.Uses[i].(*types.Func)

				if !isFunction || f.Pkg() != p.Types || f.Exported() {
					return true
				}

				o := f.Origin()
				result[o] = append(
					result[o],
					&Reference{
						Guarded:   inGuard(p, stack, wrappers),
						Enclosing: enclosing(p, stack),
					},
				)

				return true
			},
		)
	}

	return result
}
