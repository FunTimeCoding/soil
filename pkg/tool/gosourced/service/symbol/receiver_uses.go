package symbol

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func receiverUses(
	p *packages.Package,
	method *ast.FuncDecl,
	declaration *ast.Ident,
	receiver types.Object,
) []*ast.Ident {
	result := []*ast.Ident{declaration}

	if method.Body == nil {
		return result
	}

	ast.Inspect(
		method.Body,
		func(n ast.Node) bool {
			if i, okay := n.(*ast.Ident); okay && p.TypesInfo.Uses[i] == receiver {
				result = append(result, i)
			}

			return true
		},
	)

	return result
}
