package unclosed_resource

import "go/ast"

func checkedValue(
	a *ast.AssignStmt,
	index int,
) ast.Expr {
	if index != 0 || len(a.Lhs) != 2 || len(a.Rhs) != 1 {
		return nil
	}

	switch v := ast.Unparen(a.Rhs[0]).(type) {
	case *ast.IndexExpr, *ast.TypeAssertExpr, *ast.CallExpr:
		return v
	}

	return nil
}
