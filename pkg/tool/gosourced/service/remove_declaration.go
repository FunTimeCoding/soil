package service

import "go/ast"

func removeDeclaration(
	file *ast.File,
	declaration ast.Decl,
	spec ast.Spec,
) {
	g, okay := declaration.(*ast.GenDecl)

	if !okay || len(g.Specs) == 1 {
		removeFileDeclaration(file, declaration)

		return
	}

	for i, s := range g.Specs {
		if s == spec {
			g.Specs = append(g.Specs[:i], g.Specs[i+1:]...)

			return
		}
	}
}
