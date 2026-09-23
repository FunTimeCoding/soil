package expected_first

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/assert_call"
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

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || !assert_call.HasAssertPrefix(f.Name.Name) {
				continue
			}

			leading := 0

			if takesTestingFirst(p, f) {
				leading = 1
			} else if f.Recv == nil {
				continue
			}

			index := expectedIndex(f.Type.Params)

			if index < 0 || index == leading {
				continue
			}

			results.AddConcern(
				concern.NewPosition(
					"expected_first",
					fmt.Sprintf(
						"assert helper %s: expected parameter must lead",
						f.Name.Name,
					),
					p.Fset.Position(f.Pos()),
				),
			)
		}
	}
}
