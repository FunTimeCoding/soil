package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"go/ast"
)

func testNames(
	file *ast.File,
	test bool,
) map[*ast.Ident]bool {
	result := map[*ast.Ident]bool{}

	if !test {
		return result
	}

	for _, d := range file.Decls {
		if f, okay := d.(*ast.FuncDecl); okay && lint.IsTestFunction(f) {
			result[f.Name] = true
		}
	}

	return result
}
