package unread_receiver

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
	faces *face.Set,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Recv == nil || f.Body == nil || readsReceiver(p, f) {
				continue
			}

			if o := p.TypesInfo.Defs[f.Name]; o == nil || faces.Implements(o) {
				continue
			}

			results.AddConcern(
				concern.NewPosition(
					"unread_receiver",
					fmt.Sprintf(
						"%s never reads its receiver - a function, or a method of what it reads (see doc/ai/spec/naming.md, Prepositions)",
						f.Name.Name,
					),
					p.Fset.Position(f.Name.Pos()),
				),
			)
		}
	}
}
