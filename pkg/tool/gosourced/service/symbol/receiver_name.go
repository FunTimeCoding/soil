package symbol

import "go/ast"

func receiverName(f *ast.FuncDecl) string {
	if f.Recv == nil || len(f.Recv.List) == 0 {
		return ""
	}

	t := f.Recv.List[0].Type

	if s, okay := t.(*ast.StarExpr); okay {
		t = s.X
	}

	switch v := t.(type) {
	case *ast.IndexExpr:
		t = v.X
	case *ast.IndexListExpr:
		t = v.X
	}

	if i, okay := t.(*ast.Ident); okay {
		return i.Name
	}

	return ""
}
