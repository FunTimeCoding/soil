package anonymous_struct

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

		named := collectNamedStructPositions(file)
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				s, okay := n.(*ast.StructType)

				if !okay {
					return true
				}

				if named[s.Pos()] {
					return true
				}

				checkStruct(p, results, s)

				return true
			},
		)
	}
}
