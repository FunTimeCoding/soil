package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func storesClosure(
	p *packages.Package,
	a *ast.AssignStmt,
	o *types.Var,
	closures map[*types.Var]bool,
) bool {
	for i, l := range a.Lhs {
		if _, local := l.(*ast.Ident); local || i >= len(a.Rhs) {
			continue
		}

		if holdsClosure(p, a.Rhs[i], o, closures) {
			return true
		}
	}

	return false
}
