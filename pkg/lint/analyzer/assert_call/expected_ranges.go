package assert_call

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

// ExpectedRanges collects the source range of the expected
// argument - the first after *testing.T - of every assert call
// in the file. Literals inside these ranges pin actual values -
// the sanctioned home of literals - even when nested in
// constructor calls building the expected value.
func ExpectedRanges(
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

			if len(call.Args) > 2 {
				a := call.Args[1]
				result = append(result, Range{From: a.Pos(), To: a.End()})
			}

			return true
		},
	)

	return result
}
