package assert_call

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

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

			index := ExpectedIndex(p, call)

			if len(call.Args) > index+1 {
				a := call.Args[index]
				result = append(result, Range{From: a.Pos(), To: a.End()})
			}

			return true
		},
	)

	return result
}
