package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func renameReferences(
	decorations *decoration.Set,
	all []*packages.Package,
	set *token.FileSet,
	p *plannedRename,
	r *output.Results,
) error {
	for _, f := range p.References {
		owner, file := findOwningFile(all, f.Ident.Pos())

		if file == nil {
			continue
		}

		if _, e := decorations.DecorateFile(set, owner, file); e != nil {
			return e
		}

		d := decorations.DecoratedIdent(owner, f.Ident)

		if d == nil {
			continue
		}

		d.Name = p.Rename.NewName
		position := set.Position(f.Ident.Pos())
		r.AddConcern(
			concern.NewLine(
				constant.ConcernRenamed,
				fmt.Sprintf("%s → %s", p.Rename.OldName, p.Rename.NewName),
				position.Filename,
				position.Line,
				"",
				true,
			),
		)
	}

	return nil
}
