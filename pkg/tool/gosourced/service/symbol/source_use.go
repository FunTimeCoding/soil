package symbol

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strconv"
)

func sourceUse(
	p *packages.Package,
	moved map[types.Object]bool,
	packagePath string,
) (bool, bool) {
	usesMoved := false
	keepsSource := false

	for _, file := range p.Syntax {
		for _, spec := range file.Imports {
			path, f := strconv.Unquote(spec.Path.Value)

			if f == nil && path == packagePath && spec.Name != nil &&
				(spec.Name.Name == "_" || spec.Name.Name == ".") {
				keepsSource = true
			}
		}

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.SelectorExpr:
					i, okay := x.X.(*ast.Ident)

					if !okay {
						return true
					}

					name, okay := p.TypesInfo.Uses[i].(*types.PkgName)

					if okay && name.Imported().Path() == packagePath &&
						!moved[p.TypesInfo.Uses[x.Sel]] {
						keepsSource = true
					}
				case *ast.Ident:
					if moved[p.TypesInfo.Uses[x]] {
						usesMoved = true
					}
				}

				return true
			},
		)
	}

	return usesMoved, keepsSource
}
