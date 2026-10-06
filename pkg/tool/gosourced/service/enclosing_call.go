package service

import "go/ast"

func enclosingCall(path []ast.Node) *ast.CallExpr {
	if len(path) < 2 {
		return nil
	}

	callee := path[0]
	i := 1

	if s, okay := path[i].(*ast.SelectorExpr); okay && s.Sel == callee {
		callee = s
		i++
	}

	for i < len(path) {
		switch n := path[i].(type) {
		case *ast.IndexExpr:
			if n.X != callee {
				return nil
			}
		case *ast.IndexListExpr:
			if n.X != callee {
				return nil
			}
		case *ast.ParenExpr:
		default:
			c, okay := n.(*ast.CallExpr)

			if !okay || ast.Unparen(c.Fun) != ast.Unparen(callee.(ast.Expr)) {
				return nil
			}

			return c
		}

		callee = path[i]
		i++
	}

	return nil
}
