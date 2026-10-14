package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func spends(
	p *packages.Package,
	f *ast.FuncDecl,
	o *types.Func,
) bool {
	results := o.Signature().Results()

	if results.Len() == 0 {
		return false
	}

	found := false
	spent := true
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.ReturnStmt:
				if len(v.Results) == 1 && results.Len() > 1 {
					c, isCall := ast.Unparen(v.Results[0]).(*ast.CallExpr)
					spent = isCall && calleeOf(p, c) != nil
					found = true

					return spent
				}

				if len(v.Results) != results.Len() {
					spent = false

					return false
				}

				for i, e := range v.Results {
					if !mayCarry(results.At(i).Type()) || isZeroValue(p, f.Body, e) {
						continue
					}

					found = true

					if !spentValue(p, f.Body, e) {
						spent = false
					}
				}
			}

			return spent
		},
	)

	return found && spent
}
