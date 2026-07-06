package forbidden_import

import (
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

		for _, i := range file.Imports {
			path := strings.Trim(i.Path.Value, `"`)

			for key, b := range banned {
				if b.substring {
					if !strings.Contains(path, key) {
						continue
					}
				} else {
					if path != key {
						continue
					}
				}

				results.AddConcern(
					concern.NewFile(
						"forbidden_import",
						b.message,
						p.Fset.Position(i.Pos()).Filename,
						false,
					),
				)
			}
		}
	}
}
