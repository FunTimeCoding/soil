package assert_call

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func ArgumentRanges(
	p *packages.Package,
	file *ast.File,
) []Range {
	var result []Range
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			call, okay := n.(*ast.CallExpr)

			if !okay || !IsAssertCall(p, call) {
				return true
			}

			for _, argument := range call.Args {
				result = append(
					result,
					Range{From: argument.Pos(), To: argument.End()},
				)
			}

			return true
		},
	)

	return result
}
