package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func closureVariables(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
) map[*types.Var]bool {
	result := map[*types.Var]bool{}
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			a, okay := n.(*ast.AssignStmt)

			if !okay || len(a.Lhs) != len(a.Rhs) {
				return true
			}

			for i, l := range a.Lhs {
				identifier, isIdent := l.(*ast.Ident)
				f, isClosure := ast.Unparen(a.Rhs[i]).(*ast.FuncLit)

				if !isIdent || !isClosure || !references(p, f.Body, o) {
					continue
				}

				if v, okay := p.TypesInfo.ObjectOf(identifier).(*types.Var); okay {
					result[v] = true
				}
			}

			return true
		},
	)

	return result
}
