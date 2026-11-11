package pattern

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/match"
	"go/ast"
)

func expressionCall(
	p *match.Pattern,
	querySymbol string,
) *ast.CallExpr {
	statement, okay := p.Statement.(*ast.ExprStmt)

	if !okay {
		return nil
	}

	call, okay := statement.X.(*ast.CallExpr)

	if !okay {
		return nil
	}

	switch fun := call.Fun.(type) {
	case *ast.Ident:
		if fun.Name == querySymbol {
			return call
		}
	case *ast.SelectorExpr:
		if fun.Sel.Name == querySymbol {
			return call
		}
	}

	return nil
}
