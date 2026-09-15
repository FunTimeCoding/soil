package service

import (
	"go/ast"
	"go/token"
)

func findTypeDeclaration(
	file *ast.File,
	name string,
) (*ast.GenDecl, int, bool) {
	for i, d := range file.Decls {
		g, okay := d.(*ast.GenDecl)

		if !okay || g.Tok != token.TYPE {
			continue
		}

		for _, s := range g.Specs {
			t := s.(*ast.TypeSpec)

			if t.Name.Name == name {
				return g, i, len(g.Specs) > 1
			}
		}
	}

	return nil, -1, false
}
