package file_identity

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkTestFile(
	p *packages.Package,
	results *output.Results,
	file *ast.File,
	name string,
) {
	for _, d := range file.Decls {
		f, okay := d.(*ast.FuncDecl)

		if !okay || isTestFunction(f) {
			continue
		}

		results.AddConcern(
			concern.NewPosition(
				"file_identity",
				fmt.Sprintf(
					"helper %s shares %s; extract it to its own .go file",
					f.Name.Name,
					name,
				),
				p.Fset.Position(f.Pos()),
			),
		)
	}
}
