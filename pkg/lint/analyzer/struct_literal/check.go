package struct_literal

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	if p.Module == nil {
		return
	}

	module := p.Module.Path

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				if n == nil {
					return true
				}

				switch node := n.(type) {
				case *ast.UnaryExpr:
					if node.Op == token.AND {
						checkUnary(p, results, node, module)
					}
				case *ast.CallExpr:
					checkNew(p, results, node, module)
				}

				return true
			},
		)
	}
}
