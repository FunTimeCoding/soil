package symbol

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func receiverCollision(
	p *packages.Package,
	method *ast.FuncDecl,
	receiver types.Object,
	name string,
) *ast.Ident {
	var result *ast.Ident
	ast.Inspect(
		method,
		func(n ast.Node) bool {
			i, okay := n.(*ast.Ident)

			if result != nil || !okay || i.Name != name {
				return result == nil
			}

			o := p.TypesInfo.Defs[i]

			if o == nil {
				o = p.TypesInfo.Uses[i]
			}

			if o == nil || o == receiver || ownNamespace(o) {
				return true
			}

			result = i

			return false
		},
	)

	return result
}
