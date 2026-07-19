package assert_call

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

// ArgumentRanges collects the source ranges of every argument
// passed to an assert call in the file. Literals inside these
// ranges are expected values - the sanctioned home of literals.
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
