package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func returnedVariables(
	p *packages.Package,
	f *ast.FuncDecl,
) map[types.Object]bool {
	result := map[types.Object]bool{}
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.ReturnStmt:
				for _, r := range v.Results {
					if i, okay := ast.Unparen(r).(*ast.Ident); okay {
						result[p.TypesInfo.Uses[i]] = true
					}
				}
			}

			return true
		},
	)

	return result
}
