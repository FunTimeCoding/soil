package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func returnedCalls(
	p *packages.Package,
	body *ast.BlockStmt,
) map[*ast.CallExpr]bool {
	result := map[*ast.CallExpr]bool{}
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.ReturnStmt:
				for _, r := range v.Results {
					switch e := ast.Unparen(r).(type) {
					case *ast.CallExpr:
						result[e] = true
					case *ast.Ident:
						if o, okay := p.TypesInfo.Uses[e].(*types.Var); okay {
							for _, a := range assignedValues(p, body, o) {
								if c, okay := ast.Unparen(a).(*ast.CallExpr); okay {
									result[c] = true
								}
							}
						}
					}
				}
			}

			return true
		},
	)

	return result
}
