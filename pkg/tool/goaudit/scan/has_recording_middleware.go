package scan

import "go/ast"

func hasRecordingMiddleware(a ast.Expr) bool {
	l, okay := a.(*ast.CompositeLit)

	if !okay {
		return false
	}

	for _, element := range l.Elts {
		c, okay := element.(*ast.CallExpr)

		if !okay {
			continue
		}

		switch f := c.Fun.(type) {
		case *ast.IndexExpr:
			if m, okay := f.X.(*ast.SelectorExpr); okay &&
				m.Sel.Name == "RecordingMiddleware" {
				return true
			}
		case *ast.SelectorExpr:
			if f.Sel.Name == "RecordingMiddleware" {
				return true
			}
		}
	}

	return false
}
