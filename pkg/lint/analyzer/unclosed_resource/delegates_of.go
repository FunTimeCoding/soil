package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func delegatesOf(
	p *packages.Package,
	body *ast.BlockStmt,
) []*types.Func {
	var result []*types.Func
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.ReturnStmt:
				for _, e := range v.Results {
					if c, okay := e.(*ast.CallExpr); okay {
						if f := calleeOf(p, c); f != nil {
							result = append(result, f)
						}
					}
				}
			case *ast.AssignStmt:
				if f := flowingDelegate(p, body, v); f != nil {
					result = append(result, f)
				}
			}

			return true
		},
	)

	return result
}
