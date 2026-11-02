package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func assertedRelease(
	p *packages.Package,
	body *ast.BlockStmt,
	a *ast.AssignStmt,
	r *types.Var,
) []fact.Release {
	if len(a.Rhs) != 1 || len(a.Lhs) == 0 {
		return nil
	}

	t, okay := ast.Unparen(a.Rhs[0]).(*ast.TypeAssertExpr)

	if !okay {
		return nil
	}

	field := fieldOfReceiver(p, t.X, r)
	i, okay := a.Lhs[0].(*ast.Ident)

	if field == "" || !okay {
		return nil
	}

	asserted, okay := p.TypesInfo.ObjectOf(i).(*types.Var)

	if !okay || !closedVariable(p, body, asserted) {
		return nil
	}

	return []fact.Release{{Field: field}}
}
