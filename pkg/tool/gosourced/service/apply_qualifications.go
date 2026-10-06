package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/token"
	"sort"
)

func applyQualifications(
	r *output.Results,
	decorations *decoration.Set,
	set *token.FileSet,
	qualifications map[string]*relocation.FileQualification,
	targetPackagePath string,
) error {
	var filenames []string

	for filename := range qualifications {
		filenames = append(filenames, filename)
	}

	sort.Strings(filenames)

	for _, filename := range filenames {
		q := qualifications[filename]
		file, e := decorations.DecorateFile(set, q.Owner, q.File)

		if e != nil {
			return e
		}

		for ident, newName := range q.Idents {
			d := decorations.DecoratedIdent(q.Owner, ident)

			if d == nil {
				continue
			}

			d.Name = newName
			d.Path = targetPackagePath
		}

		if q.Name != nil && q.Name.Alias != "" && !q.Name.Imported {
			decorations.AddAlias(file, targetPackagePath, q.Name.Alias)
		}

		for _, qp := range q.Positions {
			r.AddConcern(
				concern.NewLine(
					"qualified",
					fmt.Sprintf(
						"%s → %s.%s",
						qp.OldName,
						q.Name.Local,
						qp.NewName,
					),
					qp.Position.Filename,
					qp.Position.Line,
					"",
					true,
				),
			)
		}
	}

	return nil
}
