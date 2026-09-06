package scan

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

func evaluatePattern(
	e ast.Expr,
	constants map[string]string,
) (string, bool) {
	switch x := e.(type) {
	case *ast.BasicLit:
		if x.Kind != token.STRING {
			return "", false
		}

		value, f := strconv.Unquote(x.Value)

		if f != nil {
			return "", false
		}

		return stripMethod(value), true
	case *ast.SelectorExpr:
		value, okay := constants[x.Sel.Name]

		if !okay {
			return "", false
		}

		return stripMethod(value), true
	case *ast.CallExpr:
		m, okay := x.Fun.(*ast.SelectorExpr)

		if !okay {
			return "", false
		}

		switch m.Sel.Name {
		case "Get", "Post", "Put", "Delete":
		default:
			return "", false
		}

		var parts []string

		for _, argument := range x.Args {
			part, okay := evaluatePattern(argument, constants)

			if !okay {
				return "", false
			}

			parts = append(parts, part)
		}

		return strings.Join(parts, ""), true
	}

	return "", false
}
