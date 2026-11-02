package unclosed_resource

import "go/ast"

func literalFields(e ast.Expr) map[string]ast.Expr {
	e = ast.Unparen(e)

	if u, okay := e.(*ast.UnaryExpr); okay {
		e = u.X
	}

	l, okay := e.(*ast.CompositeLit)

	if !okay {
		return nil
	}

	result := map[string]ast.Expr{}

	for _, element := range l.Elts {
		v, okay := element.(*ast.KeyValueExpr)

		if !okay {
			continue
		}

		if key, okay := v.Key.(*ast.Ident); okay {
			result[key.Name] = v.Value
		}
	}

	return result
}
