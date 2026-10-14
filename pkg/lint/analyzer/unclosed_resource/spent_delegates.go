package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func spentDelegates(
	p *packages.Package,
	f *ast.FuncDecl,
	o *types.Func,
) []string {
	results := o.Signature().Results()
	var result []string
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.ReturnStmt:
				for i, e := range v.Results {
					c, isCall := ast.Unparen(e).(*ast.CallExpr)

					if !isCall {
						continue
					}

					if len(v.Results) == results.Len() &&
						!carriesObligation(results.At(i).Type()) {
						continue
					}

					if g := calleeOf(p, c); g != nil {
						result = append(result, functionKey(g))
					}
				}
			}

			return true
		},
	)

	return result
}
