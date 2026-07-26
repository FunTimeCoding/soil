package service

import (
	"go/ast"
	"go/token"
)

func removeDeclaration(
	set *token.FileSet,
	file *ast.File,
	declaration ast.Decl,
	spec ast.Spec,
) {
	g, okay := declaration.(*ast.GenDecl)

	if !okay || len(g.Specs) == 1 {
		start := declaration.Pos()

		if d := declarationDocument(declaration); d != nil {
			start = d.Pos()
		}

		if spec != nil {
			d := specDocument(declaration, spec)

			if d != nil && d.Pos() < start {
				start = d.Pos()
			}
		}

		scrubComments(
			file,
			start,
			trailingCommentEnd(set, file, declaration.End()),
		)
		removeFileDeclaration(file, declaration)

		return
	}

	for i, s := range g.Specs {
		if s != spec {
			continue
		}

		start := s.Pos()

		if d := specDocument(declaration, spec); d != nil {
			start = d.Pos()
		}

		scrubComments(file, start, specSliceEnd(spec))
		g.Specs = append(g.Specs[:i], g.Specs[i+1:]...)

		return
	}
}
