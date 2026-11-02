package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"unicode"
)

func (s *Service) checkRename(
	directory string,
	all []*packages.Package,
	set *token.FileSet,
	packagePath string,
	n *symbol_rename.Rename,
	r *output.Results,
) *plannedRename {
	declaration, p, e := findDeclaration(
		all,
		packagePath,
		n.OldName,
		n.Receiver,
	)

	if e != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, e.Error(), "", false),
		)

		return nil
	}

	if v, okay := declaration.(*types.Var); okay && v.Embedded() {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
				fmt.Sprintf(
					"%s is an embedded field of %s - rename the type instead",
					n.OldName,
					n.Receiver,
				),
				"",
				false,
			),
		)

		return nil
	}

	if f := checkCollision(p, n.NewName, n.Receiver); f != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, f.Error(), "", false),
		)

		return nil
	}

	before := len(r.Entries)
	s.checkRenamedMethod(directory, all, set, declaration, r)

	if len(r.Entries) > before {
		return nil
	}

	references := resolve.FindAllReferences(all, declaration)

	if n.Receiver == "" {
		if taken := checkRenameScope(
			all,
			set,
			p.PkgPath,
			n.NewName,
			references,
		); taken != "" {
			r.AddConcern(
				concern.NewFile(constant.ConcernValidation, taken, "", false),
			)

			return nil
		}
	}

	if unicode.IsUpper(rune(n.OldName[0])) &&
		unicode.IsLower(rune(n.NewName[0])) {
		for _, f := range references {
			if f.Package.PkgPath == p.PkgPath {
				continue
			}

			position := set.Position(f.Ident.Pos())
			r.AddConcern(
				concern.NewLine(
					"cross-package",
					fmt.Sprintf(
						"%s.%s would lose access",
						f.Package.PkgPath,
						n.OldName,
					),
					position.Filename,
					position.Line,
					"",
					false,
				),
			)
		}
	}

	return &plannedRename{
		Rename:      n,
		Declaration: declaration,
		References:  references,
	}
}
