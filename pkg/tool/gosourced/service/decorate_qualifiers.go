package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"go/token"
)

func decorateQualifiers(
	r *output.Results,
	decorations *decoration.Set,
	set *token.FileSet,
	qualifiers []resolve.Reference,
	oldName string,
	newName string,
) error {
	for _, f := range qualifiers {
		position := set.Position(f.Ident.Pos())
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
		file := findSyntaxFile(set, f.Package, position.Filename)

		if file == nil {
			continue
		}

		if _, e := decorations.DecorateFile(set, f.Package, file); e != nil {
			return e
		}
	}

	return nil
}
