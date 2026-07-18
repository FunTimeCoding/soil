package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/token"
)

func renameQualifiers(
	r *output.Results,
	s *token.FileSet,
	qualifiers []resolve.Reference,
	oldName string,
	newName string,
	modified map[string]*ast.File,
) {
	for _, f := range qualifiers {
		p := s.Position(f.Ident.Pos())
		f.Ident.Name = newName
		filename := p.Filename
		i := findSyntaxFile(s, f.Package, filename)

		if i != nil {
			modified[filename] = i
		}

		r.AddConcern(
			concern.NewLine(
				"renamed",
				fmt.Sprintf("%s → %s", oldName, newName),
				p.Filename,
				p.Line,
				"",
				true,
			),
		)
	}
}
