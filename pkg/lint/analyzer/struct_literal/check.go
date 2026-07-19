package struct_literal

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/assert_call"
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

		expected := assert_call.ArgumentRanges(p, file)
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				if n == nil {
					return true
				}

				switch node := n.(type) {
				case *ast.UnaryExpr:
					if node.Op == token.AND &&
						!inExpectedArgument(expected, node.Pos()) {
						checkUnary(p, results, node, module)
					}
				case *ast.CallExpr:
					if !inExpectedArgument(expected, node.Pos()) {
						checkNew(p, results, node, module)
					}
				}

				return true
			},
		)
	}
}
