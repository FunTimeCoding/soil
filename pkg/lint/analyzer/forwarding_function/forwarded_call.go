package forwarding_function

import "go/ast"

func forwardedCall(f *ast.FuncDecl) *ast.CallExpr {
	if f.Body == nil || len(f.Body.List) != 1 {
		return nil
	}

	switch s := f.Body.List[0].(type) {
	case *ast.ReturnStmt:
		if len(s.Results) != 1 {
			return nil
		}

		call, okay := s.Results[0].(*ast.CallExpr)

		if !okay {
			return nil
		}

		return call
	case *ast.ExprStmt:
		call, okay := s.X.(*ast.CallExpr)

		if !okay {
			return nil
		}

		return call
	}

	return nil
}
