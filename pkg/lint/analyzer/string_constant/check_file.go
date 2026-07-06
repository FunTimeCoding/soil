package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkFile(
	p *packages.Package,
	results *output.Results,
	file *ast.File,
	constants map[string]knownConstant,
) {
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.CallExpr:
				checkCallArguments(p, results, v, constants)
			case *ast.IndexExpr:
				checkArgument(p, results, v.Index, constants)
			}

			return true
		},
	)
}
