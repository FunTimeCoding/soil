package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"unicode"
)

func (s *Service) renameThrough(
	directory string,
	packagePath string,
	oldName string,
	newName string,
	receiver string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := s.referenceLoad(
		directory,
		packagePath,
		"",
		func(loaded []*packages.Package, _ *token.FileSet) []types.Object {
			d, _, f := findDeclaration(loaded, packagePath, oldName, receiver)

			if f != nil {
				return nil
			}

			return []types.Object{d}
		},
	)

	if e != nil {
		return nil, e
	}

	declaration, p, e := findDeclaration(all, packagePath, oldName, receiver)

	if e != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, e.Error(), "", false),
		)

		return r, nil
	}

	if v, okay := declaration.(*types.Var); okay && v.Embedded() {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
				fmt.Sprintf(
					"%s is an embedded field of %s - rename the type instead",
					oldName,
					receiver,
				),
				"",
				false,
			),
		)

		return r, nil
	}

	e = checkCollision(p, newName, receiver)

	if e != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, e.Error(), "", false),
		)

		return r, nil
	}

	s.checkRenamedMethod(directory, all, set, declaration, r)

	if hasUnfixed(r) {
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

	e = restoreDecorations(decorations, resolve.NewNames(all), nil, out)

	if e != nil {
		return nil, e
	}

	return r, nil
}
