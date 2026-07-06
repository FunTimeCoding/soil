package defer_close

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
				d, okay := n.(*ast.DeferStmt)

				if !okay {
					return true
				}

				checkDefer(p, results, d)

				return true
			},
		)
	}
}
