package scan

import "go/ast"

func calledPackage(n ast.Node) (string, bool) {
	c, okay := n.(*ast.CallExpr)

	if !okay {
		return "", false
	}

	m, okay := c.Fun.(*ast.SelectorExpr)

	if !okay {
		return "", false
	}

	i, okay := m.X.(*ast.Ident)

	if !okay {
		return "", false
	}

	return i.Name, true
}
