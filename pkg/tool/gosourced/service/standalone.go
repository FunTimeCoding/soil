package service

import "go/ast"

func standalone(
	parent ast.Node,
	outer ast.Expr,
) bool {
	switch p := parent.(type) {
	case *ast.AssignStmt:
		return len(p.Lhs) == 1 && len(p.Rhs) == 1 && p.Rhs[0] == outer
	case *ast.ValueSpec:
		return len(p.Names) == 1 && len(p.Values) == 1 && p.Values[0] == outer
	}

	return false
}
