package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"unicode"
)

func (s *Service) Rename(
	directory string,
	packagePath string,
	oldName string,
	newName string,
	receiver string,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := resolve.LoadPackages(directory, "./...")

	if e != nil {
		return nil, e
	}

	declaration, p, e := findDeclaration(all, packagePath, oldName, receiver)

	if e != nil {
		r.AddConcern(
			concern.NewFile("validation", e.Error(), "", false),
		)

		return r, nil
	}

	e = checkCollision(p, newName, receiver)

	if e != nil {
		r.AddConcern(
			concern.NewFile("validation", e.Error(), "", false),
		)

		return r, nil
	}

	references := resolve.FindAllReferences(all, declaration)
	unexporting := unicode.IsUpper(
		rune(oldName[0]),
	) && unicode.IsLower(rune(newName[0]))

	if unexporting {
		for _, f := range references {
			if f.Package.PkgPath != p.PkgPath {
				position := set.Position(f.Ident.Pos())
				r.AddConcern(
					concern.NewLine(
						"cross-package",
						fmt.Sprintf(
							"%s.%s would lose access",
							f.Package.PkgPath,
							oldName,
						),
						position.Filename,
						position.Line,
						"",
						false,
					),
				)
			}
		}

		if hasUnfixed(r) {
			return r, nil
		}
	}

	for _, f := range references {
		position := set.Position(f.Ident.Pos())
		f.Ident.Name = newName
		r.AddConcern(
			concern.NewLine(
				"renamed",
				fmt.Sprintf("%s → %s", oldName, newName),
				position.Filename,
				position.Line,
				"",
				true,
			),
		)
	}

	_, e = writeModifiedFiles(set, references)

	if e != nil {
		return nil, e
	}

	return r, nil
}
