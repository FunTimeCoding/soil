package symbol

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func receiverShape(
	p *packages.Package,
	method string,
	declaration *ast.FuncDecl,
	receiver *types.Var,
	target *types.Named,
	moved map[string]bool,
) string {
	selected := map[*ast.Ident]bool{}
	message := ""
	ast.Inspect(
		declaration.Body,
		func(n ast.Node) bool {
			if message != "" {
				return false
			}

			s, okay := n.(*ast.SelectorExpr)

			if !okay {
				return true
			}

			i, okay := s.X.(*ast.Ident)

			if !okay || p.TypesInfo.Uses[i] != receiver {
				return true
			}

			selected[i] = true
			message = selectionShape(p, method, i.Name, s, target, moved)

			return true
		},
	)

	if message != "" {
		return message
	}

	ast.Inspect(
		declaration.Body,
		func(n ast.Node) bool {
			if i, okay := n.(*ast.Ident); okay && message == "" &&
				p.TypesInfo.Uses[i] == receiver && !selected[i] {
				message = fmt.Sprintf(
					"%s uses its receiver %s as a value - its type changes, move it by hand",
					method,
					i.Name,
				)
			}

			return message == ""
		},
	)

	return message
}
