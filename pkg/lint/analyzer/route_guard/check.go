package route_guard

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	if !inScope(p.PkgPath) {
		return
	}

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				call, okay := n.(*ast.CallExpr)

				if !okay {
					return true
				}

				checkCall(p, results, call)

				return true
			},
		)
	}
}
