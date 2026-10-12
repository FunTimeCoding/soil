package gofix

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func buildRealiasEdits(
	p *packages.Package,
	file *ast.File,
	spec *ast.ImportSpec,
	current string,
	want string,
	declaredName string,
) []Edit {
	var result []Edit

	switch {
	case want == declaredName && spec.Name != nil:
		result = append(
			result,
			Edit{position: spec.Name.Pos(), end: spec.Path.Pos(), newText: ""},
		)
	case spec.Name != nil:
		result = append(
			result,
			Edit{position: spec.Name.Pos(), end: spec.Name.End(), newText: want},
		)
	default:
		result = append(
			result,
			Edit{
				position: spec.Path.Pos(),
				end:      spec.Path.Pos(),
				newText:  fmt.Sprintf("%s ", want),
			},
		)
	}

	ast.Inspect(
		file,
		func(n ast.Node) bool {
			selector, okay := n.(*ast.SelectorExpr)

			if !okay {
				return true
			}

			ident, okay := selector.X.(*ast.Ident)

			if !okay || ident.Name != current {
				return true
			}

			if _, okay := p.TypesInfo.ObjectOf(ident).(*types.PkgName); !okay {
				return true
			}

			result = append(
				result,
				Edit{position: ident.Pos(), end: ident.End(), newText: want},
			)

			return true
		},
	)

	return result
}
