package service

import "go/ast"

func receiverName(declaration *ast.FuncDecl) string {
	if declaration.Recv == nil || len(declaration.Recv.List) == 0 {
		return ""
	}

	t := declaration.Recv.List[0].Type

	for {
		switch e := t.(type) {
		case *ast.StarExpr:
			t = e.X
		case *ast.IndexExpr:
			t = e.X
		case *ast.IndexListExpr:
			t = e.X
		case *ast.Ident:
			return e.Name
		default:
			return ""
		}
	}
}
