package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/assert_call"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkFile(
	p *packages.Package,
	results *output.Results,
	file *ast.File,
	constants map[string][]knownConstant,
) {
	expected := assert_call.ExpectedRanges(p, file)
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.CallExpr:
				if assert_call.IsAssertCall(p, v) && len(v.Args) > 2 {
					checkExpected(p, results, v.Args[1])
				}

				for _, a := range v.Args {
					checkArgument(p, results, a, constants, expected)
				}
			case *ast.IndexExpr:
				checkArgument(p, results, v.Index, constants, expected)
			}

			return true
		},
	)
}
