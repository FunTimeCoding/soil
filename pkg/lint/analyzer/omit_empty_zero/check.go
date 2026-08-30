package omit_empty_zero

import (
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

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				structure, okay := n.(*ast.StructType)

				if !okay {
					return true
				}

				for _, field := range structure.Fields.List {
					checkField(p, results, field)
				}

				return true
			},
		)
	}
}
