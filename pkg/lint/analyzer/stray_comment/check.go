package stray_comment

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		regions := emptyRegions(file)

		for _, group := range file.Comments {
			for _, c := range group.List {
				if isDirective(c.Text) {
					continue
				}

				if isEmptinessMarker(c, regions) {
					continue
				}

				results.AddConcern(
					concern.NewPosition(
						"stray_comment",
						fmt.Sprintf("comment: %s", firstLine(c.Text)),
						p.Fset.Position(c.Pos()),
					),
				)
			}
		}
	}
}
