package service

import "github.com/dave/dst"

func removeDecoratedDeclaration(
	file *dst.File,
	declaration dst.Decl,
	spec dst.Spec,
) {
	g, okay := declaration.(*dst.GenDecl)

	if !okay || spec == nil || len(g.Specs) == 1 {
		for i, d := range file.Decls {
			if d == declaration {
				file.Decls = append(file.Decls[:i], file.Decls[i+1:]...)

				return
			}
		}

		return
	}

	for i, s := range g.Specs {
		if s == spec {
			g.Specs = append(g.Specs[:i], g.Specs[i+1:]...)

			return
		}
	}
}
