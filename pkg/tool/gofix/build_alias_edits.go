package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func buildAliasEdits(
	p *packages.Package,
	file *ast.File,
	spec *ast.ImportSpec,
	d *import_decision.Decision,
) []Edit {
	current := d.Entry.Local
	want := d.Want
	var result []Edit

	switch {
	case want == d.Entry.Declared && spec.Name != nil:
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
