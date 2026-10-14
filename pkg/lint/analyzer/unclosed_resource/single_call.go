package unclosed_resource

import "go/ast"

func singleCall(f *ast.FuncLit) *ast.CallExpr {
	if f.Body == nil || len(f.Body.List) != 1 {
		return nil
	}

	statement, okay := f.Body.List[0].(*ast.ExprStmt)

	if !okay {
		return nil
	}

	result, _ := statement.X.(*ast.CallExpr)

	return result
}
