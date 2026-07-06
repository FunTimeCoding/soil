package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func buildAllEdits(
	fileSet *token.FileSet,
	all []*packages.Package,
	violations []violation,
	r *output.Results,
) []edit {
	var result []edit

	for _, v := range violations {
		path := fileSet.File(v.ident.Pos()).Name()

		if v.fix == "" {
			r.AddConcern(
				concern.NewFile(
					"collision",
					fmt.Sprintf("cannot rename %s (collision)", v.ident.Name),
					path,
					false,
				),
			)

			continue
		}

		replacement := segment.ReplaceSegment(v.ident.Name, v.segment, v.fix)
		references := resolve.FindAllReferences(all, v.object)
		r.AddConcern(
			concern.NewFile(
				"renamed",
				fmt.Sprintf(
					"renamed %s → %s (%d references)",
					v.ident.Name,
					replacement,
					len(references),
				),
				path,
				true,
			),
		)

		for _, e := range references {
			newName := segment.ReplaceSegment(e.Ident.Name, v.segment, v.fix)
			result = append(
				result,
				edit{
					position: e.Ident.Pos(),
					end:      e.Ident.End(),
					newText:  newName,
				},
			)
		}
	}

	return result
}
