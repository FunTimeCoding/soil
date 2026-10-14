package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func spentValue(
	p *packages.Package,
	body *ast.BlockStmt,
	e ast.Expr,
) bool {
	switch v := ast.Unparen(e).(type) {
	case *ast.Ident:
		o, okay := p.TypesInfo.Uses[v].(*types.Var)

		return okay && assignedOnce(p, body, o) &&
			closes(p, body, o, nil)
	case *ast.CallExpr:
		return calleeOf(p, v) != nil
	}

	return false
}
