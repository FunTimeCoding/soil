package gofix

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func buildImportAliasEdits(
	p *packages.Package,
	file *ast.File,
	spec *ast.ImportSpec,
	alias string,
	declaredName string,
) []edit {
	var result []edit
	result = append(
		result,
		edit{position: spec.Name.Pos(), end: spec.Path.Pos(), newText: ""},
	)
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			selector, okay := n.(*ast.SelectorExpr)

			if !okay {
				return true
			}

			ident, okay := selector.X.(*ast.Ident)

			if !okay {
				return true
			}

			if ident.Name != alias {
				return true
			}

			o := p.TypesInfo.ObjectOf(ident)

			if o == nil {
				return true
			}

			if _, okay := o.(*types.PkgName); !okay {
				return true
			}

			if alias != declaredName {
				result = append(
					result,
					edit{
						position: ident.Pos(),
						end:      ident.End(),
						newText:  declaredName,
					},
				)
			}

			return true
		},
	)

	return result
}
