package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"unicode"
)

func (s *Service) Rename(
	directory string,
	packagePath string,
	oldName string,
	newName string,
	receiver string,
	dryRun bool,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := resolve.LoadPackages(directory, "./...")

	if e != nil {
		return nil, e
	}

	declaration, p, e := findDeclaration(all, packagePath, oldName, receiver)

	if e != nil {
		r.AddConcern(concern.NewFile("validation", e.Error(), "", false))

		return r, nil
	}

	e = checkCollision(p, newName, receiver)

	if e != nil {
		r.AddConcern(concern.NewFile("validation", e.Error(), "", false))

		return r, nil
	}

	references := resolve.FindAllReferences(all, declaration)
	unexporting := unicode.IsUpper(rune(oldName[0])) && unicode.IsLower(
		rune(newName[0]),
	)

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

	decorations := decoration.NewSet()

	for _, f := range references {
		position := set.Position(f.Ident.Pos())
		owner, file := findOwningFile(all, f.Ident.Pos())

		if file == nil {
			continue
		}

		if _, g := decorations.DecorateFile(set, owner, file); g != nil {
			return nil, g
		}

		d := decorations.DecoratedIdent(owner, f.Ident)

		if d == nil {
			continue
		}

		d.Name = newName
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

	e = restoreDecorations(decorations, resolve.NewNames(all), nil, dryRun)

	if e != nil {
		return nil, e
	}

	if dryRun {
		r.MarkPlanned()
	}

	return r, nil
}
