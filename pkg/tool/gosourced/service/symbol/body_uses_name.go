package symbol

import "go/ast"

func bodyUsesName(
	body ast.Node,
	name string,
) bool {
	found := false
	var visit func(n ast.Node) bool
	visit = func(n ast.Node) bool {
		if found {
			return false
		}

		switch v := n.(type) {
		case *ast.SelectorExpr:
			ast.Inspect(v.X, visit)

			return false
		case *ast.KeyValueExpr:
			if _, field := v.Key.(*ast.Ident); field {
				ast.Inspect(v.Value, visit)

				return false
			}
		case *ast.Ident:
			if v.Name == name {
				found = true
			}
		}

		return true
	}
	ast.Inspect(body, visit)

	return found
}
