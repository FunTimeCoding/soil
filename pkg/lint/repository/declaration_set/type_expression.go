package declaration_set

import "go/ast"

func typeExpression(e ast.Expr) ast.Expr {
	switch v := e.(type) {
	case *ast.StarExpr:
		return typeExpression(v.X)
	case *ast.IndexExpr:
		return typeExpression(v.X)
	case *ast.IndexListExpr:
		return typeExpression(v.X)
	}

	return e
}
