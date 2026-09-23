package unclosed_resource

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func flowingDelegate(
	p *packages.Package,
	body *ast.BlockStmt,
	a *ast.AssignStmt,
) *types.Func {
	if a.Tok != token.DEFINE || len(a.Rhs) != 1 {
		return nil
	}

	call, okay := a.Rhs[0].(*ast.CallExpr)

	if !okay {
		return nil
	}

	for _, l := range a.Lhs {
		i, okay := l.(*ast.Ident)

		if !okay {
			continue
		}

		o, okay := p.TypesInfo.Defs[i].(*types.Var)

		if okay && flowsToResult(p, body, o) {
			return calleeOf(p, call)
		}
	}

	return nil
}
