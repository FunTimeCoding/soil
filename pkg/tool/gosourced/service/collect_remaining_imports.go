package service

import (
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
)

func collectRemainingImports(file *ast.File) map[string]bool {
	result := make(map[string]bool)

	for _, d := range file.Decls {
		f, okay := d.(*ast.FuncDecl)

		if okay {
			for _, spec := range imports.UsedBy(file, f) {
				result[spec.Path.Value] = true
			}

			continue
		}

		g, okay := d.(*ast.GenDecl)

		if okay {
			for _, spec := range imports.UsedBy(file, g) {
				result[spec.Path.Value] = true
			}
		}
	}

	return result
}
