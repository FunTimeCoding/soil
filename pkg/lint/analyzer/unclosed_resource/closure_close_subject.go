package unclosed_resource

import "go/ast"

func closureCloseSubject(f *ast.FuncLit) ast.Expr {
	if f.Body == nil || len(f.Body.List) != 1 {
		return nil
	}

	statement, okay := f.Body.List[0].(*ast.ExprStmt)

	if !okay {
		return nil
	}

	call, okay := statement.X.(*ast.CallExpr)

	if !okay {
		return nil
	}

	return closeSubject(call)
}
