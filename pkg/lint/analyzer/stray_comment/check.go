package stray_comment

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		name := p.Fset.File(file.Pos()).Name()

		if strings.HasSuffix(name, constant.TestSuffix) {
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
					concern.NewFile(
						"stray_comment",
						fmt.Sprintf(
							"line %d: comment: %s",
							p.Fset.Position(c.Pos()).Line,
							firstLine(c.Text),
						),
						name,
						false,
					),
				)
			}
		}
	}
}
