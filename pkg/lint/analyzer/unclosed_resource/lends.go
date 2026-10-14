package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func lends(
	p *packages.Package,
	f *ast.FuncDecl,
	o *types.Func,
	via *[]string,
) bool {
	indices := parameterIndices(o)
	results := o.Signature().Results()

	if len(indices) == 0 || results.Len() == 0 {
		return false
	}

	found := false
	lending := true
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.ReturnStmt:
				if len(v.Results) == 1 && results.Len() > 1 {
					found = true
					lending = lentFrom(
						p,
						f.Body,
						v.Results[0],
						indices,
						map[*types.Var]bool{},
						via,
					)

					return lending
				}

				if len(v.Results) != results.Len() {
					lending = false

					return false
				}

				for i, e := range v.Results {
					if !mayCarry(results.At(i).Type()) ||
						isZeroValue(p, f.Body, e) {
						continue
					}

					if !lentFrom(p, f.Body, e, indices, map[*types.Var]bool{}, via) {
						lending = false
					}

					found = true
				}
			}

			return lending
		},
	)

	return found && lending
}
